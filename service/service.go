package service

import (
	"github.com/hwb2017/CMDBDemo/dao"
	"github.com/hwb2017/CMDBDemo/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	dao *dao.Dao
}

func New(client ...*mongo.Client) *Service {
	if client != nil && len(client) == 1{
		return &Service{dao: dao.New(client[0])}
	}
	return &Service{dao: dao.New(global.MongodbClient)}
}