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
    instances := make([]ecs.Instance, 0)

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
		instances = append(instances, response.Instances.Instance...)
	}
	fmt.Println(len(instances))
    insertInstances := make([]interface{},len(instances))
    insertInstanceIds := make([]string,len(instances))

    for i, v := range instances {
    	m := make(map[string]interface{})
    	j, _ := json.Marshal(v)
    	json.Unmarshal(j, &m)
    	delete(m, "Cpu")
    	m["_id"] = m["InstanceId"]
    	insertInstances[i] = m
    	insertInstanceIds[i] = fmt.Sprintf("%v", m["_id"])
	}

	// 对比待插入实例的ID和数据库中已有实例的ID
	currentInstanceIds := make([]string,0)
	collection := global.MongodbClient.Database("infrastructure").Collection("alicloud_instance")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find().SetProjection(bson.D{{"_id", 1}})
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to insert data to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		currentInstanceIds = append(currentInstanceIds, fmt.Sprintf("%v", result["_id"]))
		if err != nil {fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)}
	}
	if err := cursor.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)
	}

	addInstances := utils.StrSliceDiff(insertInstanceIds, currentInstanceIds)
	delInstances := utils.StrSliceDiff(currentInstanceIds, insertInstanceIds)
	updateInstances := utils.StrSliceIntersection(currentInstanceIds, insertInstanceIds)
	fmt.Println(addInstances, delInstances)
	fmt.Println(len(updateInstances))


    //collection = global.MongodbClient.Database("infrastructure").Collection("alicloud_instance")
    //ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
    //defer cancel()
    //_, err = collection.InsertMany(ctx, insertInstances)
    //if err != nil {
	//	errMsg := fmt.Sprintf("Failed to insert data to mongodb: %v\n", err)
	//	fmt.Fprint(os.Stderr, errMsg)
	//	panic(errMsg)
	//}

}

func timeoutCheck(start time.Time) {
	dis := time.Since(start).Milliseconds()
	fmt.Println("elapse", dis, "ms")
}