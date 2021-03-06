package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hwb2017/CMDBDemo/lib/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response)ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (r *Response)ToResponseList(data interface{}, totalRows int) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		"pagination": Pagination{
			PageNumber: GetPageNumber(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response)ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg": err.Msg(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}