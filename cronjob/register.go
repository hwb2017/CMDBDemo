package cronjob

import (
	"github.com/robfig/cron/v3"
)

func RegisterDefaultCronJobs(c *cron.Cron) {
	c.AddFunc("@every 1m", SyncInstances)
	c.AddFunc("@every 1m", ScanAndExecuteVMLifecycle)
}