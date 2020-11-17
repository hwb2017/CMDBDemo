package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
)

func BindJSONAndValid(c *gin.Context, v interface{}) (bool, error) {
   	err := c.ShouldBindJSON(v)
   	if err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
   			return true, err
		}
		for key, value := range verrs {
			fmt.Println(key, value)
		}
		return false, verrs
	}
	return true, nil
}

func BindQueryAndValid(c *gin.Context, v interface{}) (bool, error) {
	err := c.ShouldBindQuery(v)
	if err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return true, err
		}
		for key, value := range verrs {
			fmt.Println(key, value)
		}
		return false, verrs
	}
	return true, nil
}

func Validate3yrRange(fl validator.FieldLevel) bool {
    timestamp := fl.Field().Int()
    threeYearLater := time.Now().AddDate(3,0,0).Unix()
    oneDayBefore := time.Now().AddDate(0,0,-1).Unix()
    if timestamp > oneDayBefore && timestamp < threeYearLater {
    	return true
	}
    return false
}

//func ValidateJSONDateType(field reflect.Value) interface{} {
//	if field.Type() == reflect.TypeOf(model.LocalTime{}) {
//		timeStr := field.Interface().(model.LocalTime).String()
//		if timeStr == "0001-01-01 00:00:00" {
//			return nil
//		}
//		return timeStr
//	}
//	return nil
//}