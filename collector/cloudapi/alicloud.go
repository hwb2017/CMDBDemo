package cloudapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"time"
)

func SyncAlicloudInstances() {
	defer timeoutCheck(time.Now())
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.RegionId = "cn-shanghai"
	request.PageSize = "1"

	response, err := alicloudClient.DescribeInstances(request)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get alicloud response: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	if !response.IsSuccess() {
		errMsg := fmt.Sprintf("Failed to get alicloud response\n")
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
    totalCount := response.TotalCount
    pageCount := (totalCount/100)+1
    rawInstances := make([]ecs.Instance, 0)

    for pageNumber := 1; pageNumber <=pageCount; pageNumber++ {
		request := ecs.CreateDescribeInstancesRequest()
		request.Scheme = "https"
		request.RegionId = "cn-shanghai"
		request.PageSize = "100"
		request.PageNumber = requests.Integer(strconv.Itoa(pageNumber))

		response, err := alicloudClient.DescribeInstances(request)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to get alicloud response: %v\n", err)
			fmt.Fprint(os.Stderr, errMsg)
			panic(errMsg)
		}
		if !response.IsSuccess() {
			errMsg := fmt.Sprintf("Failed to get alicloud response\n")
			fmt.Fprint(os.Stderr, errMsg)
			panic(errMsg)
		}
		// fmt.Println(response.GetHttpContentString())
		rawInstances = append(rawInstances, response.Instances.Instance...)
	}
    instancesMapping := make(map[string]interface{},len(rawInstances))
    for _, v := range rawInstances {
    	instanceAttr := make(map[string]interface{})
    	j, _ := json.Marshal(v)
    	json.Unmarshal(j, &instanceAttr)
    	delete(instanceAttr, "Cpu")
    	instanceId := fmt.Sprintf("%v", instanceAttr["InstanceId"])
    	instanceAttr["_id"] = instanceId
    	instanceAttr["_syncTime"] = time.Now().Format("2006-01-02 15:04:05")
    	instancesMapping[instanceId] = instanceAttr
	}
	instanceIds := make([]string, 0, len(instancesMapping))
	for k, _ := range instancesMapping {
		instanceIds = append(instanceIds, k)
	}

	// 对比待插入实例的ID和数据库中已有实例的ID
	storedInstanceIds := make([]string, 0, len(instancesMapping))
	collection := global.MongodbClient.Database("infrastructure").Collection("alicloud_instance")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	findOpts := options.Find().SetProjection(bson.D{{"_id", 1}})
	cursor, err := collection.Find(ctx, bson.D{}, findOpts)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to insert data to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		storedInstanceIds = append(storedInstanceIds, fmt.Sprintf("%v", result["_id"]))
		if err != nil {fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)}
	}
	if err := cursor.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)
	}

	insertInstanceIds := utils.StrSliceDiff(instanceIds, storedInstanceIds)
	deleteInstanceIds := utils.StrSliceDiff(storedInstanceIds, instanceIds)
	updateInstanceIds := utils.StrSliceIntersection(storedInstanceIds, instanceIds)

	bulkOpInstanceNum := len(insertInstanceIds) + len(deleteInstanceIds) + len(updateInstanceIds)
	bulkWriteModels := make([]mongo.WriteModel, 0, bulkOpInstanceNum)
	//设置批量操作中要插入的部分
	for _, v := range insertInstanceIds {
		instance := instancesMapping[v]
		model := mongo.NewInsertOneModel().SetDocument(instance)
		bulkWriteModels = append(bulkWriteModels, model)
	}
	// 设置批量操作中要删除的部分
	for _, v := range deleteInstanceIds {
	    model := mongo.NewDeleteOneModel().SetFilter(bson.M{"_id": v})
		bulkWriteModels = append(bulkWriteModels, model)
	}
	// 设置批量操作中要更新的部分
	for _, v := range updateInstanceIds {
	    instance := instancesMapping[v]
	    model := mongo.NewReplaceOneModel().SetFilter(bson.M{"_id": v}).SetReplacement(instance)
		bulkWriteModels = append(bulkWriteModels, model)
	}

    collection = global.MongodbClient.Database("infrastructure").Collection("alicloud_instance")
    ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
	bulkWriteOpts := options.BulkWrite().SetOrdered(false)
    res, err := collection.BulkWrite(ctx, bulkWriteModels, bulkWriteOpts)
    if err != nil {
		errMsg := fmt.Sprintf("Failed to bulk write to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	fmt.Printf("插入 %v 个，删除 %v 个，匹配 %v 个\n", res.InsertedCount, res.DeletedCount, res.MatchedCount)
}

func timeoutCheck(start time.Time) {
	dis := time.Since(start).Milliseconds()
	fmt.Println("elapse", dis, "ms")
}