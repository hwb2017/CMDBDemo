package cronjob

import "github.com/hwb2017/CMDBDemo/lib/cloudapi"

func SyncInstances() {
    cloudapi.SyncAlicloudInstances()
}
