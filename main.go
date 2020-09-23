package main

import (
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/router"
)

func init() {
	global.InitConfiguration()
	global.InitLogger()
    global.InitMongoDB()
	global.InitCronjobRunner()
}

func gracefulExit() {
    global.DisConnectMongodb()
    global.StopCronjobRunner()
}

func main() {
     defer gracefulExit()

     r := router.InitRouter()
     r.Run(global.ServerConfiguration.Host)
}