package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	defaultPageSize = 10
	maxPageSize = 50
)

type Pagination struct {
	PageNumber int
	PageSize int
	TotalRows int
}

func GetPageNumber(c *gin.Context) int {
	pageNumber, _ := strconv.Atoi(c.Query("pageNum"))
	if pageNumber <= 0 {
		return 1
	}
	return pageNumber
}

func GetPageSize(c *gin.Context) int {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
    if pageSize <= defaultPageSize {
    	return defaultPageSize
	} else if pageSize >= maxPageSize {
		return maxPageSize
	}
	return pageSize
}

func GetOffset(pageNum, pageSize int) int {
	offset := 0
	if pageNum >= 1 {
		offset = (pageNum - 1)*pageSize
	}
	return offset
}