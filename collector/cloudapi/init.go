package cloudapi

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"os"
)

var (
	alicloudAccessKey string
	alicloudAccessSecret string
	alicloudClient *ecs.Client
)

func mustGetEnv(key string) string {
	value, present := os.LookupEnv(key)
	if !present {
		errMsg := fmt.Sprintf("Environment variable %s not exists!\n", key)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	return value
}

func init() {
	alicloudAccessKey = mustGetEnv("ALICLOUD_ACCESS_KEY")
	alicloudAccessSecret = mustGetEnv("ALICLOUD_ACCESS_SECRET")
	var err error
	alicloudClient, err = ecs.NewClientWithAccessKey("cn-hangzhou", alicloudAccessKey, alicloudAccessSecret)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create alicloud client: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
}