package model

import (
	"context"
	"fmt"
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

func (v VMLifecycle) Create(client * mongo.Client) (resultID string, err error) {
	vmLifecycleCollection := VMLifecycleCollection.mongoCollection()
	result, err := vmLifecycleCollection.InsertOne(context.TODO(), v)
	if err != nil {
        return "", err
	}
	resultID = result.InsertedID.(primitive.ObjectID).String()
	return resultID, nil
}