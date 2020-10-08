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

func (d *Dao) Find(dbName, collectionName string, queryOptions *model.QueryOptions) ([]bson.M, error){
	collection := model.Collection{
		DBName: dbName,
		CollectionName: collectionName,
	}
	return collection.Find(d.client, queryOptions)
}

func (d *Dao) BulkSync(dbName, collectionName string, bulkSyncModels model.BulkSyncModels) error {
	collection := model.Collection{
		DBName: dbName,
		CollectionName: collectionName,
	}
	return collection.BulkSync(d.client, bulkSyncModels)
}