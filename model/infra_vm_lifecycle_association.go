package model

import (
	"context"
	"github.com/hwb2017/CMDBDemo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VMLifecycleAssociationCollection struct {
	Collection
}

type VMLifecycleAssociation struct {
	VMLifecycleID string `json:"vm_lifecycle_id"`
	VMID string `json:"vm_id"`
}

type vmLifecycleAssociation struct {
	VMLifecycleID primitive.ObjectID
	VMID string
}

func convertVMLifecycleAssociation(v VMLifecycleAssociation) (interface{}, error) {
    _v := vmLifecycleAssociation{}
    _v.VMID = v.VMID
    id, err := primitive.ObjectIDFromHex(v.VMLifecycleID)
    _v.VMLifecycleID = id
    return _v, err
}

func (v *VMLifecycleAssociationCollection) setup() {
	v.DBName = "infrastructure"
	v.CollectionName = "vm_lifecycle_association"
}

func (v *VMLifecycleAssociationCollection) BulkCreate(client *mongo.Client, docs []VMLifecycleAssociation) error {
	v.setup()
	collection := v.mongodbCollection(client)
	var documents []interface{}
	for _, d := range docs {
		document, err := convertVMLifecycleAssociation(d)
		if err != nil {
			return err
		}
		documents = append(documents, document)
	}
	_, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}
	return nil
}

func (v *VMLifecycleAssociationCollection) BulkUpdate(client *mongo.Client, vmLifecycleID string, vmIDs []string) error {
	v.setup()
	collection := v.mongodbCollection(client)
	storedVMLifecycleID, err := primitive.ObjectIDFromHex(vmLifecycleID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteMany(context.TODO(), bson.D{{
		"vmlifecycleid", storedVMLifecycleID,
	}})
	if err != nil {
		return err
	}
	docs := make([]vmLifecycleAssociation,0)
	for _, v := range vmIDs {
		docs = append(docs, vmLifecycleAssociation{
			VMLifecycleID: storedVMLifecycleID,
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
	storedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteMany(context.TODO(), bson.D{{
		"vmlifecycleid", storedID,
	}})
	return err
}