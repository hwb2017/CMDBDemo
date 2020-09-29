package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	DBName string
	CollectionName string
}

type BulkSyncModels struct {
	ModelMapping map[string]interface{}
	InsertIDs []string
	DeleteIDs []string
	UpdateIDs []string
}

func (c Collection) mongodbCollection(client *mongo.Client) *mongo.Collection{
	return client.Database(c.DBName).Collection(c.CollectionName)
}

func (c Collection) FindWithProjection(client *mongo.Client, fields ...string) (results []bson.M ,err error){
	collection := c.mongodbCollection(client)
	projection := make(bson.M, 0)
    for _, field := range fields {
    	if len(field) <= 0 {
    		continue
		}
        projection[field] = true
	}
	findOpts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), bson.M{}, findOpts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, err
}

// BulkSync combine insert, delete and replace data and execute then in bulk
func (c Collection) BulkSync(client *mongo.Client, m BulkSyncModels) error {
	collection := c.mongodbCollection(client)

	bulkOpsNum := len(m.InsertIDs) + len(m.DeleteIDs) + len(m.UpdateIDs)
	bulkWriteModels := make([]mongo.WriteModel, 0, bulkOpsNum)
	//设置批量同步中要插入的部分
	for _, v := range m.InsertIDs {
		doc := m.ModelMapping[v]
		model := mongo.NewInsertOneModel().SetDocument(doc)
		bulkWriteModels = append(bulkWriteModels, model)
	}
	//设置批量同步中要删除的部分
	for _, v := range m.DeleteIDs {
		model := mongo.NewDeleteOneModel().SetFilter(bson.M{"_id": v})
		bulkWriteModels = append(bulkWriteModels, model)
	}
	//设置批量同步中要更新的部分
	for _, v := range m.UpdateIDs {
		doc := m.ModelMapping[v]
		model :=  mongo.NewReplaceOneModel().SetFilter(bson.M{"_id": v}).SetReplacement(doc)
		bulkWriteModels = append(bulkWriteModels, model)
	}
	bulkWriteOpts := options.BulkWrite().SetOrdered(false)
	_, err := collection.BulkWrite(context.TODO(), bulkWriteModels, bulkWriteOpts)
	if err != nil {
		return err
	}
	return nil
}