package cronjob

import (
	cloudapi "github.com/hwb2017/CMDBDemo/collector/cloudapi"
	"github.com/hwb2017/CMDBDemo/executor"
	"github.com/robfig/cron/v3"
)

func RegisterDefaultCronJobs(c *cron.Cron) {
	c.AddFunc("@every 1m", cloudapi.SyncAlicloudInstances)
	c.AddFunc("@every 1m", executor.ScanAndExecuteVMLifecycle)
}