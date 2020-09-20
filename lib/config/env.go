package config

import (
	"fmt"
	"os"
)

type CloudApiConfig struct {
	AliCloudAccessKey    string
	AliCloudAccessSecret string
}

type EnvConfig struct {
	CloudApiConfig
}

func mustGetEnv(key string) string {
	value, present := os.LookupEnv(key)
	if !present {
		errMsg := fmt.Sprintf("Environment variable %s not exists!\n", key)
		panic(errMsg)
	}
	return value
}

func ReadEnvVars() *EnvConfig {
    envConfig := new(EnvConfig)
    envConfig.AliCloudAccessKey = mustGetEnv("ALICLOUD_ACCESS_KEY")
	envConfig.AliCloudAccessSecret = mustGetEnv("ALICLOUD_ACCESS_SECRET")
	return envConfig
}