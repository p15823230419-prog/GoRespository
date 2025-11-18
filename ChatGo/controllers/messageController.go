package controllers

import (
	"ChatGo/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendHandler 发送消息请求函数
func SendHandler(c *gin.Context) {
	var mes models.Message

	if err := c.ShouldBind(&mes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	//从token中提取id
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{
			"code": 1,
			"msg":  "未找到userId",
		})
	}
	mes.SenderId = userId.(uint64)
	//添加到数据库
	result := db.Create(&mes)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"code": 1,
			"msg":  "插入新数据错误",
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{},
	})
}

// MessagesHandler 读取所有消息API
func MessagesHandler(c *gin.Context) {

	var msgs []models.Message

	user := c.Query("user")
	afterID := c.Query("after_id")

	//gorm查询消息表
	err := db.
		Model(models.Message{}).
		Select("id", "senderId", "content", "created_at").
		Where("receiver = ? AND id > ?", user, afterID).
		Order("id ASC").
		Find(&msgs).
		Error
	if err != nil {
		log.Println("未获取到消息")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "未获取到消息",
			"data": nil,
		})
		return
	}
	log.Printf("成功获取 %d 条消息。\n", len(msgs))
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"messages": msgs,
		},
	})
}
