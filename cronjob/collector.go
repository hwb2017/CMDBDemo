package cronjob

import (
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/service"
)

func SyncInstances() {
    svc := service.New()
    err := svc.SyncAlicloudInstances()
    if err != nil {
		global.Logger.Errorf("SyncAlicloudInstances err: %v", err)
    	return
	}
	return
}
