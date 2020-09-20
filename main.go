package main

import (
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/router"
)

func init() {
	global.InitConfiguration()
    global.InitMongoDB()
}

func gracefulExit() {
    global.DisConnectMongodb()
}

func main() {
     defer gracefulExit()

     //cloudapi.SyncAlicloudInstances()
     //global.InitCronjobRunner()
     //cron.RegisterCronFuncs(global.CronjobRunner)
     //global.StartCronjobRunner()
     //defer global.StopCronjobRunner()
     //select {}

     r := router.InitRouter()
     r.Run(global.ServerConfiguration.Host)
}