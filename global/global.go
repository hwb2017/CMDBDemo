package global

import (
	"github.com/hwb2017/CMDBDemo/lib/config"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DatabaseConfiguration config.DatabaseConfig
	ServerConfiguration config.ServerConfig
	LogConfiguration config.LogConfig
	CloudApiConfiguration config.CloudApiConfig
    CronjobRunner *cron.Cron
    MongodbClient *mongo.Client
	Logger *logrus.Logger
)
