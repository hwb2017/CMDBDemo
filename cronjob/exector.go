package cronjob

import (
	"fmt"
	"github.com/hwb2017/CMDBDemo/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

func ScanAndExecuteVMLifecycle() {
	vmLifecycleCollection := global.MongodbClient.Database("infrastructure").Collection("vm_lifecycle")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	lookupStage := bson.D{
		{"$lookup",bson.D{
			{"from", "vm_lifecycle_association"},
			{"localField", "_id"},
			{"foreignField", "VMLifecycleID"},
			{"as", "associated_vm_ids"},
		},
		}}

	cursor, err := vmLifecycleCollection.Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		errMsg := fmt.Sprintf("Failed to find data in mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse mongodb document: %v\n", err)
	}
}
