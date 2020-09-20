package infrastructure

import (
	"context"
	"fmt"
	"github.com/hwb2017/CMDBDemo/global"
	"go.mongodb.org/mongo-driver/bson"
	"os"
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
}

func (v *VMLifecycle) Create() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
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
}