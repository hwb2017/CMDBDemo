package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/service"
	"net/http"
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
	svc := service.New()
	results, err := svc.ListVMLifecycle()
	if err != nil {
		global.Logger.Errorf("ListVMLifecycle err: %v", err)
		c.JSON(http.StatusOK,gin.H{
			"code": 500,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
	return
}