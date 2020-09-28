package initialize

import (
	"context"
	"fmt"
	"github.com/hwb2017/CMDBDemo/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func DisConnectMongodb() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := global.MongodbClient.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func InitMongoDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var mongodbURI string
	if global.DatabaseConfiguration.Username == "" && global.DatabaseConfiguration.Password == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s",global.DatabaseConfiguration.Host)
	} else {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s@%s/%s",
			global.DatabaseConfiguration.Username,
			global.DatabaseConfiguration.Password,
			global.DatabaseConfiguration.Host,
			global.DatabaseConfiguration.DBName)
	}
	global.MongodbClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create mongodb client: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = global.MongodbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		errMsg := fmt.Sprintf("Failed to ping mongodb instance: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
}