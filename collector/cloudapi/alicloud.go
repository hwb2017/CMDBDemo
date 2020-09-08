package cloudapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/hwb2017/CMDBDemo/global"
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

    for i, v := range instances {
    	m := make(map[string]interface{})
    	j, _ := json.Marshal(v)
    	json.Unmarshal(j, &m)
    	delete(m, "Cpu")
    	m["_id"] = m["InstanceId"]
    	insertInstances[i] = m
	}
    collection := global.MongodbClient.Database("infrastructure").Collection("alicloud_instance")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _, err = collection.InsertMany(ctx, insertInstances)
    if err != nil {
		errMsg := fmt.Sprintf("Failed to insert data to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}

}

func timeoutCheck(start time.Time) {
	dis := time.Since(start).Milliseconds()
	fmt.Println("elapse", dis, "ms")
}