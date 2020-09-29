package cronjob

import (
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/service"
)

func SyncInstances() {
    svc := service.New()
    err := svc.SyncAliCloudInstances()
    if err != nil {
		global.Logger.Errorf("SyncAliCloudInstances err: %v", err)
    	return
	}
    err = svc.SyncAWSInstances()
	if err != nil {
		global.Logger.Errorf("SyncAWSInstances err: %v", err)
		return
	}
	return
}