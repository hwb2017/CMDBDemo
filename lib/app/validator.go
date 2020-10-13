package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValid(c *gin.Context, v interface{}) error {
   	err := c.ShouldBind(v)
   	if err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
   			return nil
		}
		for key, value := range verrs {
			fmt.Println(key, value)
		}
		return verrs
	}
	return nil
}
