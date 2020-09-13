package global

import (
	"github.com/robfig/cron/v3"
)

var CronjobRunner *cron.Cron

func InitCronjobRunner() {
	CronjobRunner = cron.New()
}

func StartCronjobRunner() {
	CronjobRunner.Start()
}

func StopCronjobRunner() {
	CronjobRunner.Stop()
}