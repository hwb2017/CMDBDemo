package dao

import (
	"github.com/hwb2017/CMDBDemo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dao struct {
	client *mongo.Client
}

func New(client *mongo.Client) *Dao {
	return &Dao{client: client}
}

func (d *Dao) FindWithProjection(dbName, collectionName string, fields ...string) ([]bson.M, error){
	collection := model.Collection{
		DBName: dbName,
		CollectionName: collectionName,
	}
	return collection.FindWithProjection(d.client, fields...)
}

func (d *Dao) BulkSync(dbName, collectionName string, bulkSyncModels model.BulkSyncModels) error {
	collection := model.Collection{
		DBName: dbName,
		CollectionName: collectionName,
	}
	return collection.BulkSync(d.client, bulkSyncModels)
}