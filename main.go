package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"message": "pong",
		//})
		c.String(http.StatusOK,"pong")
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") //前种方式的缩写
		c.String(http.StatusOK, "Hello %s %s ", firstname, lastname)
	})
	//获取Post参数
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") //设置默认值
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	//Get Post 结合
	//如：POST /post?id=1234&page=1 HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded
	//name=manu&message=this_is_great
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		logs.Info("id:%s; page:%s;name:%s;message:%s", id, page, name, message)
	})
	//上传文件
	//给表单限制上传大小 默认32MB
	router.MaxMultipartMemory = 8 << 20 //8MB
	router.POST("/upload", func(c *gin.Context) {
		//单文件
		file, _ := c.FormFile("file")
		logs.Info(file.Filename)
		//上传文件到指定路径
		_ = c.SaveUploadedFile(file, "./files/"+file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s'uploaded", file.Filename))
	})
	router.POST("/uploadMulti", func(c *gin.Context) {
		//多文件
		form, _ := c.MultipartForm()
		files := form.File["upload"]

		for _, file := range files {
			logs.Info(file.Filename)
			_ = c.SaveUploadedFile(file, "./files/"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})

	//路由分组
	//Simple group :v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("read", readEndpoint)
	}
	// Simple Group:v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("read", readEndpoint)
	}

	//无中间件启动
	//r:=gin.New()
	//默认启动方式，包括logger、recovery中间件
	//r:=gin.Default()

	//使用中间件
	//创建一个不包含中间件的路由器
	r := gin.New()

	//全局中间件
	//使用Logger中间件
	r.Use(gin.Logger())

	//使用Recovery中间件
	r.Use(gin.Recovery())

	//路由添加中间件，可以添加任意多个
	//r.GET("/benchmark",MyBenchLogger(),benchEndpoint)
	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)
		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}
	//_ = r.Run(":8080")

	//禁用console颜色
	gin.DisableConsoleColor()

	//创建记录日志的文件
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	//日志同时写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)



	_ = router.Run(":9090")
}

func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func benchEndpoint(context *gin.Context) {

}

func analyticsEndpoint(context *gin.Context) {

}

func readEndpoint(c *gin.Context) {

}

func submitEndpoint(c *gin.Context) {

}

func loginEndpoint(c *gin.Context) {

}
