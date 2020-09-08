package main

import (
	"github.com/hwb2017/CMDBDemo/collector/cloudapi"
	"github.com/hwb2017/CMDBDemo/global"
)

func initialization() {
    global.InitMongoDB()
}

func gracefulExit() {
    global.DisConnectMongodb()
}

func main() {
     initialization()
     defer gracefulExit()
     cloudapi.SyncAlicloudInstances()
}