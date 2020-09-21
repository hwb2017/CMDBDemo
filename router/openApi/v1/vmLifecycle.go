package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"time"
)

type VMOperation int

const (
	Stop VMOperation = iota
	Destroy
)

type VMLifecycleRule struct {
	Operation VMOperation `json:"operation"`
	ActionTime string `json:"actionTime"`
}

type VMLifecycleRequest struct {
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant"`
	VMLifecycleRules []VMLifecycleRule `json:"vmLifecycleRules"`
	VMIDs []string `json:"vmIDs"`
}

// CreateVMLifecycle create vmLifecycleStrategy and association with virtual machines
func CreateVMLifecycle(c *gin.Context) {
	vmLifecycleReq := &service.CreateVMLifecycleRequest{}
    c.ShouldBindJSON(&vmLifecycleReq)
	svc := service.New()
	err := svc.CreateVMLifecycle(vmLifecycleReq)
	if err != nil {
		global.Logger.Errorf("CreateVMLifecycle err: %v", err)
		c.JSON(http.StatusOK,gin.H{
			"code": 500,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
	})
	return
}

func ListVMLifecycle(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}