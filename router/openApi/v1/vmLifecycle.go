package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"go.mongodb.org/mongo-driver/bson"
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
	ActionTime time.Time `json:"actionTime"`
}

type VMLifecycleRequest struct {
	Maintainer string `json:"maintainer"`
	Applicant string `json:"applicant"`
	VMLifecycleRules []VMLifecycleRule `json:"vmLifecycleRules"`
	VMIDs []string `json:"vmIDs"`
}

// CreateVMLifecycleStrategy create vmLifecycleStrategy and association with virtual machines
func CreateVMLifecycleStrategy(c *gin.Context) {
    // Maintainer, Applicant, Rules, VMs
	vmLifecycleReq := VMLifecycleRequest{}
    c.ShouldBindJSON(&vmLifecycleReq)

	vmLifecycleCollection := global.MongodbClient.Database("infrastructure").Collection("vm_lifecycle")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	vmLifecycle := bson.D{
		{"VMLifecycleRules", vmLifecycleReq.VMLifecycleRules},
		{"Applicant", vmLifecycleReq.Applicant},
		{"Maintainer", vmLifecycleReq.Maintainer},
		{"CreationTime", time.Now().Format("2006-01-02 15:04:05")},
		{"UpdateTime", time.Now().Format("2006-01-02 15:04:05")},
	}
	res, err := vmLifecycleCollection.InsertOne(ctx, vmLifecycle)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to insert to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}

	vmLifecycleID := res.InsertedID
	vmLifecycleAssociations := make([]interface{},0)
    for _, vmID := range vmLifecycleReq.VMIDs {
    	vmLifecycleAssociations = append(vmLifecycleAssociations, bson.D{
    		{"VMLifecycleID",vmLifecycleID},
    		{"VMID", vmID},
		})
	}
	vmLifecycleAssociationCollection := global.MongodbClient.Database("infrastructure").Collection("vm_lifecycle_association")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = vmLifecycleAssociationCollection.InsertMany(ctx, vmLifecycleAssociations)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to insert to mongodb: %v\n", err)
		fmt.Fprint(os.Stderr, errMsg)
		panic(errMsg)
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
	})
}