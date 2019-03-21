package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {
		c.String(200, "dd")
	})
	_ = http.ListenAndServe(":5454", router)

}
