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
}

func ReturnBindError(c *gin.Context, err error, data ...interface{}) {
	c.JSON(200, gin.H{
		"code": -1,
		"msg":  PareJSONError(err),
		"data": data,
	})
}

func ReturnSuccess(c *gin.Context, msg interface{}, data ...interface{}) {
	var body interface{}
	if len(data) == 1 {
		// 正常对象
		body = data[0]
	} else if len(data) == 0 {
		// 默认空对象
		body = gin.H{}
	} else {
		// 多个参数时才用数组
		body = data
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": body,
	})
	return
}

func ReturnError(c *gin.Context, err error, data ...interface{}) {
	c.JSON(200, gin.H{
		"code": -1,
		"msg":  err.Error(),
		"data": data,
	})
}
