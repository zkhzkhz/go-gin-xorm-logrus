package controller

import (
	"../log"
	orm "../xorm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHelp(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		log.Warn("id 未输入")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "id未输入",
		})
		return
	}
	v := orm.GetHelp(id)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   v,
	})
	return
}
