package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AliCloudInstanceCollection struct {
	Collection
}

type AWSInstanceCollection struct {
	Collection
}

func (v *VMCollection) setup() {
	v.Collection.DBName = "infrastructure"
	v.Collection.CollectionName = "vm_basic_view"
}

func (a *AliCloudInstanceCollection) setup() {
	a.DBName = "infrastructure"
	a.CollectionName = "alicloud_instance"
}

func (a *AWSInstanceCollection) setup() {
	a.DBName = "infrastructure"
	a.CollectionName = "aws_instance"
}

type VMCollection struct {
	Collection
	AliCloudInstanceCollection
	AWSInstanceCollection
}

func (v *VMCollection) ListBasicView(client *mongo.Client, queryOptions *QueryOptions) (interface{}, error) {
	v.setup()
	collection := v.mongodbCollection(client)
    err := v.AliCloudInstanceCollection.ListBasicView(client, &QueryOptions{})
    if err != nil {
    	return nil, err
	}
	err = v.AWSInstanceCollection.ListBasicView(client, &QueryOptions{})
	if err != nil {
		return nil, err
	}
    cursor, err := v.find(collection, queryOptions)
	return v.handleCursor(cursor)
}

func (v *VMCollection) Count(client *mongo.Client, queryOptions *QueryOptions) (int, error) {
	v.setup()
	collection := v.mongodbCollection(client)
	total, err := collection.CountDocuments(context.TODO(), queryOptions.filter)
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (a *AliCloudInstanceCollection) ListBasicView(client *mongo.Client, queryOptions *QueryOptions) error {
	a.setup()
	collection := a.mongodbCollection(client)
	projection := bson.M{
		"PublicIpAddress": bson.M{
			"$arrayElemAt": bson.A{"$PublicIpAddress.IpAddress", 0},
		},
		"InstanceName": 1,
		"OSName": 1,
		"InstanceType": 1,
		"PrivateIpAddress": bson.M{
			"$arrayElemAt": bson.A{"$NetworkInterfaces.NetworkInterface.PrimaryIpAddress",0},
		},
		"EipIpAddress": "$EipAddress.IpAddress",
	}
	queryOptions.WithProjection(projection)
	extendPipelines := make([]bson.D, 0)
    projectStage := bson.D{{
        "$project", bson.M{
			"vm_provider": "alicloud",
			"instance_id": "$_id",
			"instance_type": "$InstanceType",
			"instance_name": "$InstanceName",
			"os_name": "$OSName",
			"private_ip_address": "$PrivateIpAddress",
			"public_ip_address": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$EipIpAddress",bson.TypeNull}},
					"$EipIpAddress",
					"$EipIpAddress"},
			},
		},
	}}
	mergeStage := bson.D{{
		"$merge", bson.M{
			"into": "vm_basic_view",
			"on": "_id",
			"whenMatched": "replace",
			"whenNotMatched": "insert",
		},
	}}
	extendPipelines = append(extendPipelines, projectStage, mergeStage)
	queryOptions.WithExtendAggregationPipelineStages(extendPipelines)
	_, err := a.aggregate(collection, queryOptions)
	return err
}

func (a *AWSInstanceCollection) ListBasicView(client *mongo.Client, queryOptions *QueryOptions) error {
	a.setup()
	collection := a.mongodbCollection(client)
	projection := bson.M{
		"PublicIpAddress": 1,
		"Tags": 1,
		"Platform": 1,
		"InstanceType": 1,
		"PrivateIpAddress": 1,
	}
	queryOptions.WithProjection(projection)
	extendPipelines := make([]bson.D, 0)
	addFieldsStage := bson.D{{
		"$addFields", bson.M{
			"instanceName": bson.M{
				"$filter": bson.M{
					"input": "$Tags",
					"as": "tag",
					"cond": bson.M{
						"$eq": bson.A{"$$tag.Key", "Name"},
					},
				},
			},

		},
	}}
	unwindStage := bson.D{{
		"$unwind", "$instanceName",
	}}
	projectinStage := bson.D{{
        "$project", bson.M{
        	"vm_provider": "aws",
			"instance_id": "$_id",
			"instance_type": "$InstanceType",
			"instance_name": "$instanceName.Value",
			"os_name": bson.M{
        		"$ifNull": bson.A{"$Platform","Unknown"},
        		},
			"private_ip_address": "$PrivateIpAddress",
			"public_ip_address": "$PublicIpAddress",
        },
	}}
	mergeStage := bson.D{{
		"$merge", bson.M{
			"into": "vm_basic_view",
			"on": "_id",
			"whenMatched": "replace",
			"whenNotMatched": "insert",
		},
	}}
	extendPipelines = append(extendPipelines, addFieldsStage, unwindStage, projectinStage, mergeStage)
	queryOptions.WithExtendAggregationPipelineStages(extendPipelines)
	_, err := a.aggregate(collection, queryOptions)
	return err
}