package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	DBName string
	CollectionName string
}

type Find struct {
	Collection
	projection bson.M
}

func (c Collection) mongodbCollection(client *mongo.Client) *mongo.Collection{
	return client.Database(c.DBName).Collection(c.CollectionName)
}

func (c Collection) FindWithProjection(client *mongo.Client, fields ...string) ([]bson.M ,error){
	collection := c.mongodbCollection(client)

}