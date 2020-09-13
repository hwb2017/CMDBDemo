package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/hwb2017/CMDBDemo/router/baseApi/v1"
)

func InitRouter() *gin.Engine{
	r := gin.New()
	root := r.Group("")

	baseApi := root.Group("baseApi")
	vm := baseApi.Group("virtualMachine")
	vm.GET("ListVMBasicView", v1.ListVMBasicView)
	return r
}
