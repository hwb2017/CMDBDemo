package infrastructure

import "github.com/hwb2017/CMDBDemo/global"

var vmLifecycleCollection = global.MongodbClient.Database("infrastructure").Collection("vm_lifecycle")