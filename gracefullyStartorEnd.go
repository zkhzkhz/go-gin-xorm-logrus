package main

import (
	"./config"
	"./controller"
	"./log"
	"./models"
	_ "bufio"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lestrrat/go-file-rotatelogs"
	_ "github.com/pkg/errors"
	_ "github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	_ "path"
	"time"
)

//自定义结构体绑定表单数据

func GetDataC(c *gin.Context) {
	var b models.StructC
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"a":    b.NestedStructPointer,
		"c.go": b.FieldC,
	})
}
func GetDataD(c *gin.Context) {
	var b models.StructD
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This c.go.ShouldBind consumes c.go.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.go.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {

	}
}
func SomeHandler1(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	//c.go.ShouldBindBodyWith 在绑定之前将body存储到上下文中，这对性能有轻微影响，因此如果你要立即调用，则不应使用此方法
	//此功能仅适用于这些格式 — JSON, XML, MsgPack, ProtoBuf。对于其他格式，Query, Form, FormPost, FormMultipart,
	//可以被c.ShouldBind()多次调用而不影响性能
	// This reads c.go.Request.Body and stores the result into the context.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// At this time, it reuses body stored in the context.
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// And it can accepts other formats
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {

	}
}

func main() {

	//gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	router := gin.New()
	router.Use(log.Logger(), gin.Recovery())
	//跨域设置
	router.Use(cors.Default())
	//sessions
	//cookie based
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		_ = session.Save()
		c.JSON(http.StatusOK, gin.H{"count": count})
	})
	//redis
	//storeRedis, _ := redis.NewStore(10, "tcp", "loaclhost:6379", "", []byte("secret"))
	//router.Use(sessions.Sessions("mysession",storeRedis))
	//router.GET("/incrRedis", func(c.go *gin.Context) {
	//	session:=sessions.Default(c.go)
	//	var count int
	//	v:=session.Get("count")
	//	if v==nil {
	//		count=0
	//	}else{
	//		count=v.(int)
	//		count++
	//	}
	//	session.Set("count",count)
	//	_ = session.Save()
	//	c.go.JSON(http.StatusOK,gin.H{"count":count})
	//})
	////memcache
	//client := mc.NewMC("localhost:11211", "username", "password")
	//storeMemcache := memcached.NewMemcacheStore(client, "", []byte("secret"))
	//router.Use(sessions.Sessions("mysession", storeMemcache))
	//
	//router.GET("/incrMemcache", func(c.go *gin.Context) {
	//	session := sessions.Default(c.go)
	//	var count int
	//	v := session.Get("count")
	//	if v == nil {
	//		count = 0
	//	} else {
	//		count = v.(int)
	//		count++
	//	}
	//	session.Set("count", count)
	//	session.Save()
	//	c.go.JSON(200, gin.H{"count": count})
	//})
	////MongoDB
	//session, err := mgo.Dial("localhost:27017/test")
	//if err != nil {
	//	// handle err
	//}
	//
	//c.go := session.DB("").C("sessions")
	//storeMongo := mongo.NewStore(c.go, 3600, true, []byte("secret"))
	//router.Use(sessions.Sessions("mysession", storeMongo))
	//router.GET("/incrMongo", func(c.go *gin.Context) {
	//	session := sessions.Default(c.go)
	//	var count int
	//	v := session.Get("count")
	//	if v == nil {
	//		count = 0
	//	} else {
	//		count = v.(int)
	//		count++
	//	}
	//	session.Set("count", count)
	//	session.Save()
	//	c.go.JSON(200, gin.H{"count": count})
	//})

	//d, _ := time.ParseDuration("4354434588s")
	//log.Info(d)
	var pp *os.File
	//var err error
	defer pp.Close()
	pp, _ = os.OpenFile("./log/api.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	//log.SetLogLevel(logrus.InfoLevel)
	//log.Info(os.Stdout)
	//log.SetLogFormatter(logrus.Formatter())
	//log := logrus.New()
	mw := io.MultiWriter(os.Stdout, pp)
	//log.Out = pp
	//gin.DefaultWriter = io.MultiWriter(os.Stdout)
	logrus.SetOutput(mw)
	gin.SetMode(gin.DebugMode)
	//取配置文件信息
	projectConfig, _ := config.GetConfig()
	port := projectConfig.Channel.EmayReminderConfig.Url
	fmt.Println(port)
	//log.SetOutput(os.Stdout)
	log.Info("test")
	log.Info("msg")
	log.Info("dddd")
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	//router.GET("/getb", controller.GetDataB)
	router.GET("/getc", GetDataC)
	router.GET("/getd", GetDataD)
	router.POST("/login", controller.Signin)
	router.POST("/welcome", controller.Welcome)
	router.POST("/refresh", controller.Refresh)
	srv := &http.Server{
		Addr:    ":8050",
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Error("listen: %s\n", err)
		}
	}()
	// a timeout of 5 seconds.
	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}
