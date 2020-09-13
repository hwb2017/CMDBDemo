package cron

import (
	"github.com/hwb2017/CMDBDemo/collector/cloudapi"
	"github.com/robfig/cron/v3"
)

func RegisterCronFuncs(c *cron.Cron) {
	c.AddFunc("@every 1m", cloudapi.SyncAlicloudInstances)
}