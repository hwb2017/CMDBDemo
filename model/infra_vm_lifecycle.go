package model

import (
	"context"
	"fmt"
	"github.com/hwb2017/CMDBDemo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type VMOperation int64

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

func TranslateVMOperation(v int64) (op string, err error) {
	vmOperation := VMOperation(v)
	switch vmOperation {
	case StopOperation:
		return "stop", nil
	case DestroyOperation:
		return "destroy", nil
	default:
		return "", fmt.Errorf("%v not a valid vm operation", v)
	}
}

type VMLifecycleCollection struct {
	Collection
}

type VMLifecycleRule struct {
	Operation VMOperation `json:"operation"`
	ActionTime int64 `json:"action_time"`
}

type VMLifecycle struct {
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant"`
	VMLifecycleRules []VMLifecycleRule `json:"vm_lifecycle_rules"`
	VMIDs []string `json:"vm_ids"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

func (v *VMLifecycleCollection) setup() {
	v.DBName = "infrastructure"
	v.CollectionName = "vm_lifecycle"
}

func (v *VMLifecycleCollection) Create(client *mongo.Client, doc VMLifecycle) (resultID string, err error) {
	v.setup()
	collection := v.mongodbCollection(client)
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
        return "", err
	}
	resultID = result.InsertedID.(primitive.ObjectID).Hex()
	return resultID, nil
}

func (v *VMLifecycleCollection) Update(client *mongo.Client, id string, doc VMLifecycle) error {
	v.setup()
	collection := v.mongodbCollection(client)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{
		"$set", bson.D{
			{"maintainer",doc.Maintainer},
			{"applicant",doc.Applicant},
			{"vmlifecyclerules",doc.VMLifecycleRules},
			{"vmids",doc.VMIDs},
			{"updatetime",doc.UpdateTime},
		},
	}}
	result, err := collection.UpdateOne(context.TODO(),filter,update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("update error, unmodified vm lifecycle id: %v", id)
	}
	return nil
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
			{"as", "associated_vm_docs"},
		},
		}}
	addFieldsStage := bson.D{
		{"$addFields",bson.D{
			{
				"associated_vm_ids",bson.D{
				{"$map",bson.D{
					{"input","$associated_vm_docs"},
					{"as","vm_doc"},
					{"in","$$vm_doc.vmid"},
				}},
			},
			},
		},
		}}
	extendPipelines = append(extendPipelines, lookupStage, addFieldsStage)
	queryOptions.WithExtendAggregationPipelineStages(extendPipelines)
	cursor, err := v.aggregate(collection, queryOptions)
	if err != nil {
		return nil, err
	}
	return v.handleCursor(cursor)
}

func (v *VMLifecycleCollection) GetWithAssociation(client *mongo.Client, id string) (interface{}, error) {
	v.setup()
	collection := v.mongodbCollection(client)
	extendPipelines := make([]bson.D, 0)
	lookupStage := bson.D{
		{"$lookup",bson.D{
			{"from", "vm_lifecycle_association"},
			{"localField", "_id"},
			{"foreignField", "VMLifecycleID"},
			{"as", "associated_vm_docs"},
		},
		}}
	addFieldsStage := bson.D{
		{"$addFields",bson.D{
			{
				"associated_vm_ids",bson.D{
				{"$map",bson.D{
					{"input","$associated_vm_docs"},
					{"as","vm_doc"},
					{"in","$$vm_doc.vmid"},
				}},
			},
			},
		},
		}}
	extendPipelines = append(extendPipelines, lookupStage, addFieldsStage)
	queryOptions := &QueryOptions{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	queryOptions.WithFilter(bson.M{"_id": objectId})
	queryOptions.WithExtendAggregationPipelineStages(extendPipelines)
	cursor, err := v.aggregate(collection, queryOptions)
	if err != nil {
		return nil, err
	}
	return v.handleCursor(cursor)
}

func (v *VMLifecycleCollection) Delete(client *mongo.Client, id string) (int, error) {
	v.setup()
	collection := v.mongodbCollection(client)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	res, err := collection.DeleteOne(context.TODO(), bson.D{{
		"_id", objectId,
	}})
	return int(res.DeletedCount), err
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

func (v *VMLifecycleCollection) handleCursor(cursor *mongo.Cursor) ([]bson.M, error) {
	defer cursor.Close(context.TODO())

	results := make([]bson.M, 0, 10)
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {return nil, err}
		vmLifecycleRules := utils.InterfaceSlice(result["vmlifecyclerules"])
		for _, v := range vmLifecycleRules {
            vmLifecycleRule := v.(bson.M)
			vmLifecycleRule["operation"], err = TranslateVMOperation(vmLifecycleRule["operation"].(int64))
			if err != nil {
				return nil, err
			}
		}
		result["vmlifecyclerules"] = vmLifecycleRules
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}