package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

// 绑定为json
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	f, _ := os.Create("./bind/bind.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	gin.DefaultWriter = io.MultiWriter(f)
	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	// Example for binding XML (
	//    <?xml version="1.0" encoding="UTF-8"?>
	//    <root>
	//        <user>user</user>
	//        <password>123</password>
	//    </root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you have logined"})
	})

	router.GET("/someAsciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Serves unicode entities
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// Serves literal characters
	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	//重定向
	//外部链接
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	//Gin内部路由重定向，使用HandleContext
	router.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	// Listen and serve on 0.0.0.0:8080
	_ = router.Run(":8080")
}

//若要将请求主体绑定到结构体中，请使用模型绑定，目前支持JSON、XML、YAML和标准表单值(foo=bar&boo=baz)的绑定。
//
//Gin使用 go-playground/validator.v8 验证参数，查看完整文档。
//
//需要在绑定的字段上设置tag，比如，绑定格式为json，需要这样设置 json:"fieldname" 。
//
//此外，Gin还提供了两套绑定方法：
//
//Must bind
//Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML
//Behavior - 这些方法底层使用 MustBindWith，如果存在绑定错误，请求将被以下指令中止 c.AbortWithError(400, err).SetType(ErrorTypeBind)，响应状态代码会被设置为400，请求头Content-Type被设置为text/plain; charset=utf-8。注意，如果你试图在此之后设置响应代码，将会发出一个警告 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422，如果你希望更好地控制行为，请使用ShouldBind相关的方法
//Should bind
//Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML
//Behavior - 这些方法底层使用 ShouldBindWith，如果存在绑定错误，则返回错误，开发人员可以正确处理请求和错误。
//当我们使用绑定方法时，Gin会根据Content-Type推断出使用哪种绑定器，如果你确定你绑定的是什么，你可以使用MustBindWith或者BindingWith。
//
//你还可以给字段指定特定规则的修饰符，如果一个字段用binding:"required"修饰，并且在绑定时该字段的值为空，那么将返回一个错误
