package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
	"time"
)

type Booking struct {
	CheckIn time.Time `form: "check_in" binding: "required, bookabledate" time_format: "2006-01-02"`
	CheckOut time.Time `form: "check_out" binding: "required, gtfield-CheckIn" time_format: "2006-01-02"`
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,) bool {
    if date, ok := field.Interface().(time.Time); ok {
    	today := time.Now()
    	if today.Year() > date.Year() || today.YearDay() > date.YearDay(){
    		return false
		}
	}
    return false
}

func main(){
	route := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookableDate", bookableDate)
	}
	route.GET("/bookable", getBookable)
	route.Run(":8080")
}

func getBookable(c *gin.Context)  {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil{
		c.JSON(200, gin.H{
			"message": "Booking dates are valid!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
