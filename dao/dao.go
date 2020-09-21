package dao

import "go.mongodb.org/mongo-driver/mongo"

type Dao struct {
	client *mongo.Client
}

func New(client *mongo.Client) *Dao {
	return &Dao{client: client}
}