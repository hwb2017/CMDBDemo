package cronjob

import (
	"github.com/robfig/cron/v3"
)

func RegisterDefaultCronJobs(c *cron.Cron) {
	c.AddFunc("0 0 * * *", SyncInstances)
	c.AddFunc("0 0 * * *", ScanAndExecuteVMLifecycle)
}