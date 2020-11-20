package cloudapi

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/hwb2017/CMDBDemo/utils"
	"strconv"
)

func (a *AliCloud) DescribeInstances() ([]interface{} ,error) {
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.RegionId = "cn-shanghai"
	request.PageSize = "1"

	response, err := a.Client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	if !response.IsSuccess() {
		return nil, errors.New("DescribeInstances response not success")
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

		response, err := a.Client.DescribeInstances(request)
		if err != nil {
			return nil, err
		}
		if !response.IsSuccess() {
			return nil, err
		}
		rawInstances = append(rawInstances, response.Instances.Instance...)
	}
	return utils.InterfaceSlice(rawInstances), nil
}

func (a *AliCloud) StopInstances(ids []string) error {
	request := ecs.CreateStopInstancesRequest()
	request.Scheme = "https"
	request.InstanceId = &ids

	response, err := a.Client.StopInstances(request)
	if err != nil {
		return err
	}
	if !response.IsSuccess() {
		return errors.New("StopInstance response not success")
	}
	return nil
}