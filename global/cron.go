package global

import (
	"github.com/hwb2017/CMDBDemo/cronjob"
	"github.com/robfig/cron/v3"
)

var CronjobRunner *cron.Cron

func InitCronjobRunner() {
	CronjobRunner = cron.New()
	cronjob.RegisterDefaultCronJobs(CronjobRunner)
	CronjobRunner.Start()
}

func StopCronjobRunner() {
	CronjobRunner.Stop()
}