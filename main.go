package main

import (
	"github.com/hwb2017/CMDBDemo/cronjob"
	"github.com/hwb2017/CMDBDemo/global/initialize"
)

func init() {
	initialize.InitConfiguration()
	initialize.InitLogger()
    initialize.InitMongoDB()
	initialize.InitCronjobRunner()
}

func gracefulExit() {
    initialize.DisConnectMongodb()
    initialize.StopCronjobRunner()
}

func main() {
     defer gracefulExit()

     cronjob.SyncInstances()
     //r := router.InitRouter()
     //r.Run(global.ServerConfiguration.Host)
}