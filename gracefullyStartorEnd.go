package main

import (
	_ "bufio"
	"context"
	"fmt"
	"gin/config"
	. "gin/log"
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
type StructA struct {
	FieldA string `form:"field_a"`
}
type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}
type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	_ = c.Bind(&b)
	Info(b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
func GetDataC(c *gin.Context) {
	var b StructC
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}
func GetDataD(c *gin.Context) {
	var b StructD
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
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {

	}
}
func SomeHandler1(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	//c.ShouldBindBodyWith 在绑定之前将body存储到上下文中，这对性能有轻微影响，因此如果你要立即调用，则不应使用此方法
	//此功能仅适用于这些格式 — JSON, XML, MsgPack, ProtoBuf。对于其他格式，Query, Form, FormPost, FormMultipart,
	//可以被c.ShouldBind()多次调用而不影响性能
	// This reads c.Request.Body and stores the result into the context.
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
	router.Use(Logger(), gin.Recovery())
	//d, _ := time.ParseDuration("4354434588s")
	//log.Info(d)
	var pp *os.File
	//var err error
	defer pp.Close()
	pp, _ = os.OpenFile("api.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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
	Info("test")
	Info("msg")
	Info("dddd")
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
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
	router.GET("/getb", GetDataB)
	router.GET("/getc", GetDataC)
	router.GET("/getd", GetDataD)
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Error("Server Shutdown:", err)
	}
	Info("Server exiting")
}
