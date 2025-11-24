package controllers

import (
	"ChatGo/models"
	"log"
	"strconv"

	"ChatGo/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// SendHandler 发送消息请求函数
func SendHandler(c *gin.Context) {
	var message models.Message

	if err := c.ShouldBind(&message); err != nil {
		utils.ReturnJSON(c, 400, utils.PareJSONError(err))
		return
	}
	//从params中提取targetId
	//从token中提取id
	userId, _ := c.Get("userId")
	message.SenderId = userId.(uint64)
	targetId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	message.ReceiverId = targetId
	if err := DB.First(&models.User{}, "id = ?", targetId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ReturnJSON(c, 1, "未找到用户")
			return
		}
	}
	//添加到数据库
	if err := DB.Create(&message).Error; err != nil {
		utils.ReturnJSON(c, 400, "创建新消息失败")
		return
	}

	utils.ReturnJSON(c, 0, "发送成功")
}

// GetMessages 读取所有消息API
func GetMessages(c *gin.Context) {

	var messages []models.Message

	userId, _ := c.Get("userId")
	var req = models.MsgListReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ReturnJSON(c, 400, utils.PareJSONError(err))
		return
	}
	//获取自己ID
	//获取聊天对象ID
	//gorm查询消息表
	if req.AfterId == 0 {
		err := DB.
			Model(models.Message{}).
			Select("id", "senderId", "receiverId", "content", "createdAt").
			Where("((receiverId = ? AND senderId = ?) or (receiverId = ? AND senderId = ?)) AND id > ?", userId, req.TargetId, req.TargetId, userId, req.AfterId).
			Order("id DESC").
			Limit(req.Limit + 1).
			Find(&messages).
			Error
		if err != nil {
			log.Println("未获取到消息")
			utils.ReturnJSON(c, 400, "未获取到消息")
			return
		}
	} else {
		err := DB.
			Model(models.Message{}).
			Select("id", "senderId", "receiverId", "content", "createdAt").
			Where("((receiverId = ? AND senderId = ?) or (receiverId = ? AND senderId = ?)) AND id < ?", userId, req.TargetId, req.TargetId, userId, req.AfterId).
			Order("id DESC").
			Limit(req.Limit + 1).
			Find(&messages).
			Error
		if err != nil {
			log.Println("未获取到消息")
			utils.ReturnJSON(c, 0, "未获取到消息")
			return
		}
	}
	messages = messages[:req.Limit]
	log.Printf("成功获取 %d 条消息。\n", len(messages))
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"next_after_id": messages[len(messages)-1].Id,
			"messages":      messages,
			"has_more":      len(messages) > req.Limit,
		},
	})
}
