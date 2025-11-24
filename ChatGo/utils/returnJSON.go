package utils

import (
	"github.com/gin-gonic/gin"
)

func ReturnJSON(c *gin.Context, code int, msg interface{}, data ...interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}
