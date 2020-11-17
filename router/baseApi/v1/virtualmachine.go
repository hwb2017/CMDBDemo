package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/lib/app"
	"github.com/hwb2017/CMDBDemo/lib/errcode"
	"github.com/hwb2017/CMDBDemo/model"
	"github.com/hwb2017/CMDBDemo/service"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type ListVMBasicViewReq struct {
	Provider string `form:"provider"`
	IpAddr string `form:"ipAddr"`
	HostName string `form:"hostName"`
	PageNum int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

type GetVMBasicViewReq struct {
	ID string `form:"id"`
}

func ListVMBasicView(c * gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	var listVMBasicViewReq ListVMBasicViewReq
	err := c.ShouldBindQuery(&listVMBasicViewReq)
	if err != nil {
		global.Logger.Errorf("Get query string from url err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMBasicView)
		return
	}
	queryOptions := &model.QueryOptions{}
	queryOptions.WithLimit(app.GetPageSize(c))
	pageOffset := app.GetOffset(app.GetPageNumber(c), app.GetPageSize(c))
	queryOptions.WithSkip(pageOffset)
	filter := make(bson.M, 3)
	if listVMBasicViewReq.Provider != "" {
		providers := strings.Split(listVMBasicViewReq.Provider, ",")
		filter["vm_provider"] = bson.M{"$in": providers}
 	}
 	if listVMBasicViewReq.HostName != "" {
 		filter["instance_name"] = bson.M{"$regex": listVMBasicViewReq.HostName}
	}
	if listVMBasicViewReq.IpAddr != "" {
		filter["$or"] = bson.A{
			bson.M{"private_ip_address": bson.M{
				"$regex":listVMBasicViewReq.IpAddr,
			}},
			bson.M{"public_ip_address": bson.M{
				"$regex":listVMBasicViewReq.IpAddr,
			}},
		}
	}
	queryOptions.WithFilter(filter)
	vms, err := svc.ListVMBasicView(queryOptions)
	if err != nil {
		global.Logger.Errorf("ListVMBasicView err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMBasicView)
		return
	}
	totalRows, err := svc.CountVM(queryOptions)
	if err != nil {
		global.Logger.Errorf("CountVM err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMBasicView)
		return
	}
	response.ToResponseList(vms, totalRows)
	return
}

func GetVMBasicView(c * gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	var getVMBasicViewReq GetVMBasicViewReq
	err := c.ShouldBindQuery(&getVMBasicViewReq)
	if err != nil {
		global.Logger.Errorf("Get query string from url err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMBasicView)
		return
	}
	queryOptions := &model.QueryOptions{}
	filter := make(bson.M, 1)
	filter["_id"] = getVMBasicViewReq.ID
	queryOptions.WithFilter(filter)
	vm, err := svc.ListVMBasicView(queryOptions)
	if err != nil {
		global.Logger.Errorf("GetVMBasicView err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMBasicView)
		return
	}
	response.ToResponse(vm)
	return
}