package router

import (
	"github.com/gin-gonic/gin"
	bv1 "github.com/hwb2017/CMDBDemo/router/baseApi/v1"
	ov1 "github.com/hwb2017/CMDBDemo/router/openApi/v1"
)

func InitRouter() *gin.Engine{
	r := gin.New()
	root := r.Group("")

	baseApi := root.Group("baseApi")
	vm := baseApi.Group("virtualMachine")
	vm.GET("ListVMBasicView", bv1.ListVMBasicView)

	openApi := root.Group("openApi")
	vmLifecycle := openApi.Group("vmLifecycle")
	vmLifecycle.POST("CreateVMLifecycle", ov1.CreateVMLifecycleStrategy)
	return r
}
