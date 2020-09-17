package global

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

var MongodbClient *mongo.Client

func DisConnectMongodb() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := MongodbClient.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func InitMongoDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongodbClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create mongodb client: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = MongodbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		errMsg := fmt.Sprintf("Failed to ping mongodb instance: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
}