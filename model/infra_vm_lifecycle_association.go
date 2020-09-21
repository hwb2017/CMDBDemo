package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type VMLifecycleAssociationCollection struct {}

type VMLifecycleAssociation struct {
	VMLifecycleID string
	VMID string
}

func (v VMLifecycleAssociationCollection) mongodbCollection(client * mongo.Client) *mongo.Collection{
	return client.Database("infrastructure").Collection("vm_lifecycle")
}

func (v VMLifecycleAssociationCollection) BulkCreate(client *mongo.Client, docs []VMLifecycleAssociation) error {
	vmLifecycleAssociationCollection := v.mongodbCollection(client)
	_, err := vmLifecycleAssociationCollection.InsertMany(context.TODO(), docs)
	if err != nil {
		return err
	}
	return nil
}