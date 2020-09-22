package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/global"
	"github.com/hwb2017/CMDBDemo/service"
	"net/http"
)

func ListVMBasicView(c * gin.Context) {
	svc := service.New()
	vms, err := svc.ListBasicView()
	if err != nil {
		global.Logger.Errorf("ListVMBasicView err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": vms,
	})
	return
}