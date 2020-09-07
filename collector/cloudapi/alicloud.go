package cloudapi

import "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

func syncAlicloudInstances() {
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou",ak,asecurity)

	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	response, err := client.DescribeInstances(request)
	if err != nil {

	}
}