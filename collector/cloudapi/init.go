package cloudapi

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/hwb2017/CMDBDemo/global"
	"os"
)

var (
	alicloudClient *ecs.Client
)

func init() {
	alicloudAccessKey := global.CloudApiConfiguration.AliCloudAccessKey
	alicloudAccessSecret := global.CloudApiConfiguration.AliCloudAccessSecret
	var err error
	alicloudClient, err = ecs.NewClientWithAccessKey("cn-hangzhou", alicloudAccessKey, alicloudAccessSecret)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create alicloud client: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
}