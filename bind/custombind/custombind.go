package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = v.RegisterValidation("bookabledate", bookableDate)
	}
	router.GET("/bookable", getBookable)

	//router.Any("/testing", startPage)
	//https://github.com/julienschmidt/httprouter/issues/12
	router.GET("/user/:name/:id", startPage1)
	_ = router.Run(":8085")
}

func startPage1(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// 如果是Get，那么接收不到请求中的Post的数据？？
	// 如果是Post, 首先判断 `content-type` 的类型 `JSON` or `XML`, 然后使用对应的绑定器获取数据.
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	//if c.ShouldBind(&person) == nil {
	//	log.Println("====== Only Bind By Query String ======")
	//	//	log.Println(person.ID)
	//	log.Println(person.Name)
	//	log.Println(person.Address)
	//	log.Println(person.Birthday)
	//}
	if err := c.ShouldBindUri(&person); err != nil {
		logs.Warn(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	c.String(http.StatusOK, "Success", person.Birthday)
}

type Person struct {
	ID       string    `uri:"id" binding:"required,uuid"`
	Name     string    `form:"name" binding:"required" uri:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-01" time_utc:"1"`
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Booking dates are valid!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
