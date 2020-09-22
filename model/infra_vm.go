package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type vmBasicView struct  {
	VMProvider string `json:"vm_provider"`
	InstanceID string `json:"instance_id"`
	InstanceType string `json:"instance_type"`
	InstanceName string `json:"instance_name"`
	OSName string `json:"os_name"`
	PublicIpAddress string `json:"public_ip_addresses"`
	PrivateIpAddress string `json:"private_ip_addresses"`
}

type AlicloudInstanceCollection struct {}

func (a AlicloudInstanceCollection) mongodbCollection(client * mongo.Client) *mongo.Collection{
	return client.Database("infrastructure").Collection("alicloud_instance")
}

type VMCollection struct{
	AlicloudInstanceCollection
}

func (v VMCollection) ListBasicView(client *mongo.Client) (vms []vmBasicView, err error){
	vms = make([]vmBasicView, 0)
    alicloudVMs, err := v.AlicloudInstanceCollection.ListBasicView(client)
    if err != nil {
    	return nil, err
	}
	vms = append(vms, alicloudVMs...)
	return vms, err
}

func (a AlicloudInstanceCollection) ListBasicView(client *mongo.Client) (vms []vmBasicView, err error) {
	collection := a.mongodbCollection(client)
	projectionStage := bson.D{
		{"$project",bson.D{
			{"PublicIpAddress", bson.D{
				{"$arrayElemAt", bson.A{
					"$PublicIpAddress.IpAddress", 0},
				},
			}},
			{"InstanceName",1},
			{"OSName",1},
			{"InstanceType",1},
			{"PrivateIpAddress",bson.D{
				{"$arrayElemAt",bson.A{
					"$NetworkInterfaces.NetworkInterface.PrimaryIpAddress",0,
				}},
			}},
			{"EipIpAddress","$EipAddress.IpAddress"},
		},
		}}
	cursor, err := collection.Aggregate(context.TODO(),mongo.Pipeline{projectionStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	vms = make([]vmBasicView, 0)
	for cursor.Next(context.TODO()) {
		var result bson.M
		var publicIpAddress string
		err := cursor.Decode(&result)
		if v, ok := result["PublicIpAddress"]; ok {
			publicIpAddress = v.(string)
		} else {
			publicIpAddress = result["EipIpAddress"].(string)
		}
		vm := vmBasicView{
			VMProvider: "alicloud",
			InstanceID: result["_id"].(string),
			InstanceType: result["InstanceType"].(string),
			InstanceName: result["InstanceName"].(string),
			OSName: result["OSName"].(string),
			PrivateIpAddress: result["PrivateIpAddress"].(string),
			PublicIpAddress: publicIpAddress,
		}
		vms = append(vms, vm)
		if err != nil {return nil, err}
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return vms, err
}