package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/lib/app"
	"github.com/hwb2017/CMDBDemo/lib/errcode"
	"github.com/hwb2017/CMDBDemo/model"
	"github.com/hwb2017/CMDBDemo/service"
)

// CreateVMLifecycle create vmLifecycleStrategy and association with virtual machines
func CreateVMLifecycle(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	vmLifecycleReq := &service.CreateVMLifecycleRequest{}
	valid, err := app.BindJSONAndValid(c, vmLifecycleReq)
    if !valid {
		global.Logger.Errorf("Invalid param err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	} else if err != nil {
		global.Logger.Errorf("Parse param err: %v", err)
		response.ToErrorResponse(errcode.ParseParamsError)
		return
	}
	err = svc.CreateVMLifecycle(vmLifecycleReq)
	if err != nil {
		global.Logger.Errorf("CreateVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateVMLifecycle)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func UpdateVMLifecycle(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	vmLifecycleReq := &service.UpdateVMLifecycleRequest{}
	valid, err := app.BindJSONAndValid(c, vmLifecycleReq)
	if !valid {
		global.Logger.Errorf("Invalid param err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	} else if err != nil {
		global.Logger.Errorf("Parse param err: %v", err)
		response.ToErrorResponse(errcode.ParseParamsError)
		return
	}
	err = svc.UpdateVMLifecycle(vmLifecycleReq)
	if err != nil {
		global.Logger.Errorf("UpdateVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateVMLifecycle)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func ListVMLifecycle(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	queryOptions := &model.QueryOptions{}
	queryOptions.WithLimit(app.GetPageSize(c))
	pageOffset := app.GetOffset(app.GetPageNumber(c), app.GetPageSize(c))
	queryOptions.WithSkip(pageOffset)
	results, err := svc.ListVMLifecycle(queryOptions)
	if err != nil {
		global.Logger.Errorf("ListVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMLifecycle)
		return
	}
	totalRows, err := svc.CountVMLifecycle(queryOptions)
	if err != nil {
		global.Logger.Errorf("CountVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMLifecycle)
		return
	}
	response.ToResponseList(results, totalRows)
	return
}

func GetVMLifecycle(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	vmLifecycleReq := &service.GetVMLifecycleRequest{}
	valid, err := app.BindQueryAndValid(c, vmLifecycleReq)
	if !valid {
		global.Logger.Errorf("Invalid param err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	result, err := svc.GetVMLifecycle(vmLifecycleReq.VMLifecycleID)
	if err != nil {
		global.Logger.Errorf("GetVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetVMLifecycle)
		return
	}
	response.ToResponse(result)
	return
}

func DeleteVMLifecycle(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	vmLifecycleReq := &service.DeleteVMLifecycleRequest{}
	valid, err := app.BindQueryAndValid(c, vmLifecycleReq)
	if !valid {
		global.Logger.Errorf("Invalid param err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	deletedCount, err := svc.DeleteVMLifecycle(vmLifecycleReq.VMLifecycleID)
	if err != nil {
		global.Logger.Errorf("DeleteVMLifecycle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteVMLifecycle)
		return
	}
	response.ToResponse(gin.H{
		"deleted_count": deletedCount,
	})
	return
}