package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

type VMOperation uint32

const (
	StopOperation VMOperation = iota
	DestroyOperation
)

func ParseVMOperation(op string) (VMOperation, error) {
	switch strings.ToLower(op) {
	case "stop":
		return StopOperation, nil
	case "destroy":
		return DestroyOperation, nil
	}
	var vmOp VMOperation
	return vmOp, fmt.Errorf("not a valid vm operation")
}

type VMLifecycleCollection struct {}

type VMLifecycleRule struct {
	Operation VMOperation `json:"operation"`
	ActionTime time.Time `json:"action_time"`
}

type VMLifecycle struct {
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant"`
	VMLifecycleRules []VMLifecycleRule `json:"vm_lifecycle_rules"`
	VMIDs []string `json:"vm_ids"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (v VMLifecycleCollection) mongodbCollection(client * mongo.Client) *mongo.Collection{
	return client.Database("infrastructure").Collection("vm_lifecycle")
}

func (v VMLifecycleCollection) Create(client * mongo.Client, doc VMLifecycle) (resultID string, err error) {
	collection := v.mongodbCollection(client)
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
        return "", err
	}
	resultID = result.InsertedID.(primitive.ObjectID).String()
	return resultID, nil
}

func (v VMLifecycleCollection) ListWithAssociation(client *mongo.Client) (results []bson.M,err error) {
	collection := v.mongodbCollection(client)
	lookupStage := bson.D{
		{"$lookup",bson.D{
			{"from", "vm_lifecycle_association"},
			{"localField", "_id"},
			{"foreignField", "VMLifecycleID"},
			{"as", "associated_vm_ids"},
		},
		}}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{lookupStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {return nil, err}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}