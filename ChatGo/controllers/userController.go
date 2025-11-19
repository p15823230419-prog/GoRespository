package controllers

import (
	"ChatGo/models"
	"ChatGo/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// RegisterUser 注册请求post
func RegisterUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(500, gin.H{
			"code": 20001,
			"msg":  "参数错误" + err.Error(),
		})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		c.JSON(500, gin.H{
			"code": 20001,
			"msg":  "格式错误" + err.Error(),
		})
		return
	}
	//密码加密
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{"msg": "密码加密失败"})
		return
	}
	req.Password = hashedPassword
	//注册用户
	if err := db.Create(&req).Error; err != nil {
		c.JSON(401, gin.H{
			"code": 1,
			"msg":  "用户名已注册" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"msg":    "注册成功",
			"userId": req.Id,
		},
	})
}

// DeleteUser 删除用户DELETE
func DeleteUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "参数错误" + err.Error(),
		})
		return
	}
	if err := db.Delete(&req).Error; err != nil {
		c.JSON(401, gin.H{
			"code": 1,
			"msg":  "删除失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "删除成功",
	})
}

// UpdateUser 更改用户数据put
func UpdateUser(c *gin.Context) {
	var req models.User
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "参数错误",
			"data": err.Error(),
		})
		return
	}
	if req.Password != "" {
		//密码加密
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			c.JSON(500, gin.H{"msg": "密码加密失败"})
			return
		}
		req.Password = hashedPassword
	}
	//获取当前用户id
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{
			"code": 1,
			"msg":  "未找到userId",
		})
		return
	}
	userIdInt, _ := userId.(uint64)
	req.Id = userIdInt
	//更新数据库
	if err := db.Model(&req).Updates(req).Error; err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "未找到此id的用户",
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "更新成功",
	})
}

// LoginUser 登录请求函数post
func LoginUser(c *gin.Context) {
	var req models.User

	// 1. 获取 JSON
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": "参数错误",
		})
		return
	}

	// 2. 查询用户
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{
			"code": 1,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	//3. 验证密码
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(401, gin.H{
			"code": 1,
			"msg":  "密码错误",
		})
		return
	}

	//生成token
	token, _ := utils.GenerateToken(user.Id, user.Username)
	// 4. 成功
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"token": token,
		},
	})
}

// SelectUser 查询用户:输入查询的值进行查询
func SelectUser(c *gin.Context) {
	id := c.Query("id")
	username := c.Query("username")
	if id != "" {
		var user models.User
		if err := db.
			Select("id", "username", "avatar", "nickname", "createdAt", "updatedAt").
			Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(401, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": user,
		})
		return
	} else if username != "" {
		var req []models.User
		if err := db.
			Select("id", "username", "avatar", "nickname", "createdAt", "updatedAt").
			Where("username LIKE ?", "%"+username+"%").Find(&req).Error; err != nil {
			c.JSON(500, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": req,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "请提供要查询的值",
		})
	}
}
