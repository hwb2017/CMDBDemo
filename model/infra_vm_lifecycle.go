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
	_ VMOperation = iota
	StopOperation
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
	return vmOp, fmt.Errorf("%s not a valid vm operation", op)
}

type VMLifecycleCollection struct {
	Collection
}

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

func (v *VMLifecycleCollection) setup() {
	v.DBName = "infrastructure"
	v.CollectionName = "vm_lifecycle"
}

func (v *VMLifecycleCollection) Create(client * mongo.Client, doc VMLifecycle) (resultID string, err error) {
	v.setup()
	collection := v.mongodbCollection(client)
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
        return "", err
	}
	resultID = result.InsertedID.(primitive.ObjectID).String()
	return resultID, nil
}

func (v *VMLifecycleCollection) ListWithAssociation(client *mongo.Client, queryOptions *QueryOptions) (interface{}, error) {
	v.setup()
	collection := v.mongodbCollection(client)
	extendPipelines := make([]bson.D, 0)
	lookupStage := bson.D{
		{"$lookup",bson.D{
			{"from", "vm_lifecycle_association"},
			{"localField", "_id"},
			{"foreignField", "VMLifecycleID"},
			{"as", "associated_vm_ids"},
		},
		}}
	extendPipelines = append(extendPipelines, lookupStage)
	queryOptions.WithExtendAggregationPipelineStages(extendPipelines)
	cursor, err := v.aggregate(collection, queryOptions)
	if err != nil {
		return nil, err
	}
	return v.handleCursor(cursor)
}

func (v *VMLifecycleCollection) Count(client *mongo.Client, queryOptions *QueryOptions) (int, error) {
	v.setup()
	collection := v.mongodbCollection(client)
	total, err := collection.CountDocuments(context.TODO(), queryOptions.filter)
	if err != nil {
		return 0, err
	}
	return int(total), nil
}