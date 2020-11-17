package model

import (
	"context"
	"github.com/hwb2017/CMDBDemo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VMLifecycleAssociationCollection struct {
	Collection
}

type VMLifecycleAssociation struct {
	VMLifecycleID string `json:"vm_lifecycle_id"`
	VMID string `json:"vm_id"`
}

func (v *VMLifecycleAssociationCollection) setup() {
	v.DBName = "infrastructure"
	v.CollectionName = "vm_lifecycle_association"
}

func (v *VMLifecycleAssociationCollection) BulkCreate(client *mongo.Client, docs []VMLifecycleAssociation) error {
	v.setup()
	collection := v.mongodbCollection(client)
	documents := utils.InterfaceSlice(docs)
	_, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}
	return nil
}

func (v *VMLifecycleAssociationCollection) BulkUpdate(client *mongo.Client, vmLifecycleID string, vmIDs []string) error {
	v.setup()
	collection := v.mongodbCollection(client)
	_, err := collection.DeleteMany(context.TODO(), bson.D{{
		"vmlifecycleid", vmLifecycleID,
	}})
	if err != nil {
		return err
	}
	docs := make([]VMLifecycleAssociation,0)
	for _, v := range vmIDs {
		docs = append(docs, VMLifecycleAssociation{
			VMLifecycleID: vmLifecycleID,
			VMID: v,
		})
	}
	documents := utils.InterfaceSlice(docs)
	_, err = collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}
	return nil
}

func (v *VMLifecycleAssociationCollection) DeleteByVMLifecycleID(client *mongo.Client, id string) error {
	v.setup()
	collection := v.mongodbCollection(client)
	_, err := collection.DeleteMany(context.TODO(), bson.D{{
		"vmlifecycleid", id,
	}})
	return err
}