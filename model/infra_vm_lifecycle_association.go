package model

import (
	"context"
	"github.com/hwb2017/CMDBDemo/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type VMLifecycleAssociationCollection struct {}

type VMLifecycleAssociation struct {
	VMLifecycleID string `json:"vm_lifecycle_id"`
	VMID string `json:"vm_id"`
}

func (v VMLifecycleAssociationCollection) mongodbCollection(client * mongo.Client) *mongo.Collection{
	return client.Database("infrastructure").Collection("vm_lifecycle_association")
}

func (v VMLifecycleAssociationCollection) BulkCreate(client *mongo.Client, docs []VMLifecycleAssociation) error {
	vmLifecycleAssociationCollection := v.mongodbCollection(client)
	documents := utils.InterfaceSlice(docs)
	_, err := vmLifecycleAssociationCollection.InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}
	return nil
}