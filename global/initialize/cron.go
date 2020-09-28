package initialize

import (
	"github.com/hwb2017/CMDBDemo/cronjob"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/robfig/cron/v3"
)

func InitCronjobRunner() {
	global.CronjobRunner = cron.New()
	cronjob.RegisterDefaultCronJobs(global.CronjobRunner)
	global.CronjobRunner.Start()
}

func StopCronjobRunner() {
	global.CronjobRunner.Stop()
}