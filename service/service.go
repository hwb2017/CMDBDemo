package service

import (
	"github.com/hwb2017/CMDBDemo/dao"
	"github.com/hwb2017/CMDBDemo/global"
)

type Service struct {
	dao *dao.Dao
}

func New() Service {
	return Service{dao: dao.New(global.MongodbClient)}
}