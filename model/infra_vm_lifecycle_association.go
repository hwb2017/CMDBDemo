package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type VMLifecycleAssociation struct {
	VMLifecycleID string
	VMID string
}

func (v VMLifecycleAssociation) Create(client *mongo.Client) error {
	vmLifecycleAssociationCollection := VMLifecycleAssociationCollection.mongoCollection()
	_, err := vmLifecycleAssociationCollection.InsertOne(context.TODO(), v)
	if err != nil {
        return err
	}
	return nil
}