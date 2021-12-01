package main

import (
	. "github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *Engine {
	r := Default()
	r.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	_ = r.Run(":8048")
}
