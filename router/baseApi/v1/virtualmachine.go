package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/lib/app"
	"github.com/hwb2017/CMDBDemo/lib/errcode"
	"github.com/hwb2017/CMDBDemo/service"
)

func ListVMBasicView(c * gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	vms, err := svc.ListBasicView()
	if err != nil {
		global.Logger.Errorf("ListVMBasicView err: %v", err)
		response.ToErrorResponse(errcode.ErrorListVMBasicView)
		return
	}
	response.ToResponse(vms)
	return
}