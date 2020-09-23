package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/lib/app"
	"github.com/hwb2017/CMDBDemo/lib/errcode"
	"github.com/hwb2017/CMDBDemo/service"
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
	response := app.NewResponse(c)
	svc := service.New()
	err := svc.CreateVMLifecycle(vmLifecycleReq)
	if err != nil {
		global.Logger.Errorf("CreateVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateVMLifecycle)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func ListVMLifecycle(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	results, err := svc.ListVMLifecycle()
	if err != nil {
		global.Logger.Errorf("ListVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMLifecycle)
		return
	}
	response.ToResponse(results)
	return
}