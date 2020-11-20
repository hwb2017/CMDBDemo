package cronjob

import (
	"context"
	"fmt"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/lib/cloudapi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func ScanAndExecuteVMLifecycle() {
	vmLifecycleCollection := global.MongodbClient.Database("infrastructure").Collection("vm_lifecycle")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	t := time.Now()
	startTime := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,t.Location())
	endTime := startTime.AddDate(0,0,1)
	startTimestamp := startTime.Unix()
	endTimestamp := endTime.Unix()
	lookupStage := bson.D{
		{"$lookup",bson.D{
			{"from", "vm_lifecycle_association"},
			{"localField", "_id"},
			{"foreignField", "vmlifecycleid"},
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
    projectStage := bson.D{
    	{"$project",bson.D{
            {"associated_vm_ids",1},
            {"vmlifecyclerules",bson.D{
            	{"$filter",bson.D{
            		{"input","$vmlifecyclerules"},
            		{"as","vmlifecyclerule"},
            		{"cond",bson.D{
            			{"$and",bson.A{
							bson.D{{"$gte",bson.A{"$$vmlifecyclerule.actiontime",startTimestamp}}},
							bson.D{{"$lt",bson.A{"$$vmlifecyclerule.actiontime",endTimestamp}}},
						}},
					}},
            	},
			}},
		},
		},
	}}
    matchStage := bson.D{
		{"$match", bson.D{
			{"vmlifecyclerules", bson.D{
				{"$gt", []interface{}{}},
			}},
		}},
	}
	cursor, err := vmLifecycleCollection.Aggregate(ctx, mongo.Pipeline{lookupStage,addFieldsStage,projectStage,matchStage})
	if err != nil {
		global.Logger.Errorf("Failed to find data in mongodb: %v", err)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			global.Logger.Errorf("Failed to parse mongodb document: %v", err)
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		global.Logger.Errorf("Failed to parse mongodb document: %v", err)
	}
	fmt.Printf("%+v\n",results)

	aliCloud, err := cloudapi.NewAliCloudClient(
		global.CloudApiConfiguration.AliCloudAccessKey,
		global.CloudApiConfiguration.AliCloudAccessSecret)
	if err != nil {
		global.Logger.Errorf("Failed to create alicloud client: %v", err)
	}
	for _, result := range results {
		vm_ids_str := make([]string,0)
		vm_ids := result["associated_vm_ids"].(primitive.A)
		for _, vm_id := range vm_ids {
			vm_ids_str = append(vm_ids_str, vm_id.(string))
		}
		aliCloud.StopInstances(vm_ids_str)
	}
}