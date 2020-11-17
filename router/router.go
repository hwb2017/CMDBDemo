package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hwb2017/CMDBDemo/lib/app"
	bv1 "github.com/hwb2017/CMDBDemo/router/baseApi/v1"
	ov1 "github.com/hwb2017/CMDBDemo/router/openApi/v1"
)

func InitRouter() *gin.Engine{
	r := gin.New()
	if v,ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("3yr-range", app.Validate3yrRange)
	}
	root := r.Group("")

	baseApi := root.Group("baseApi")
	vm := baseApi.Group("virtualMachine")
	vm.GET("ListVMBasicView", bv1.ListVMBasicView)
	vm.GET("GetVMBasicView", bv1.GetVMBasicView)

	openApi := root.Group("openApi")
	vmLifecycle := openApi.Group("vmLifecycle")
	vmLifecycle.GET("ListVMLifecycle", ov1.ListVMLifecycle)
	vmLifecycle.GET("GetVMLifecycle", ov1.GetVMLifecycle)
	vmLifecycle.POST("CreateVMLifecycle", ov1.CreateVMLifecycle)
	vmLifecycle.DELETE("DeleteVMLifecycle", ov1.DeleteVMLifecycle)
	vmLifecycle.POST("UpdateVMLifecycle", ov1.UpdateVMLifecycle)
	return r
}