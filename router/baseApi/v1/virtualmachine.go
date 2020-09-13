package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"time"
)

type vmBasicView struct  {
	InstanceType string `json:instance_type`
	InstanceName string `json:instance_name`
	OSName string `json:osname`
	PublicIpAddress string `json:public_ip_addresses`
	PrivateIpAddress string `json:private_ip_addresses`
}

func ListVMBasicView(c * gin.Context) {
	collection := global.MongodbClient.Database("infrastructure").Collection("alicloud_instance")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// EipAddress.IpAddress PublicIpAddress.IpAddress[0] InstanceName OSName InstanceType NetworkInterfaces.NetworkInterface[0].PrimaryIpAddress
	projectionStage := bson.D{
		{"$project",bson.D{
			{"_id",0},
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
	cursor, err := collection.Aggregate(ctx,mongo.Pipeline{projectionStage})
	if err != nil {
		errMsg := fmt.Sprintf("Failed to insert data to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	defer cursor.Close(ctx)

	virtualMachines := make([]vmBasicView, 0)
	for cursor.Next(ctx) {
		var result bson.M
		var publicIpAddress string
		err := cursor.Decode(&result)
		if v, ok := result["PublicIpAddress"]; ok {
			publicIpAddress = v.(string)
		} else {
			publicIpAddress = result["EipIpAddress"].(string)
		}
		vm := vmBasicView{
			InstanceType: result["InstanceType"].(string),
			InstanceName: result["InstanceName"].(string),
			OSName: result["OSName"].(string),
			PrivateIpAddress: result["PrivateIpAddress"].(string),
			PublicIpAddress: publicIpAddress,
		}
        virtualMachines = append(virtualMachines, vm)
		if err != nil {fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)}
	}
	if err := cursor.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": virtualMachines,
	})
}