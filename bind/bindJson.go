package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		var form LoginForm
		// in this case proper binding will be automatically selected
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{
					"status": "you have logined",
				})
			} else {
				// gin.H is a shortcut for map[string]interface{}
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "unauthorized",
				})
			}

		}
	})
	// gin.H is a shortcut for map[string]interface{}
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	router.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})
	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	router.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	//SecureJSON可以防止json劫持，如果返回的数据是数组，则会默认在返回值前加上"while(1)"
	// 可以自定义返回的json数据前缀
	router.SecureJsonPrefix(")]}',\n")
	router.GET("/someJSON1", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 将会输出:   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
	//JSONP可以跨域传输，如果参数中存在回调参数，那么返回的参数将是回调函数的形式
	router.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})

	//设置静态文件路径
	router.Static("/files", "./files")
	//router.StaticFS("/more_static","my_file_system")
	router.StaticFile("/favicon.ico", "./files/pic1.jpg")
	err := router.Run(":8086")
	if err != nil {
		logs.Error("gin failed to start", err)
		return
	}

}
