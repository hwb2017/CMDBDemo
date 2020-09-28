package main

import (
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/global/initialize"
	"github.com/hwb2017/CMDBDemo/router"
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

     r := router.InitRouter()
     r.Run(global.ServerConfiguration.Host)
}