package controllers

import (
	"ChatGo/models"
	"ChatGo/utils"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUser 注册请求post
func RegisterUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBind(&req); err != nil {
		utils.ReturnJSON(c, 1, utils.PareJSONError(err))
		return
	}
	//密码加密
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 20002,
			"msg":  "密码加密失败",
		})
		return
	}
	req.Password = hashedPassword
	//检查是否已存在用户名
	if err := DB.Model(models.User{}).Where("username = ?", req.Username).First(&req).Error; err == nil {
		fmt.Println(err)
		utils.ReturnJSON(c, 1, "用户名已存在")
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ReturnJSON(c, 1, "数据库错误")
		return
	}
	//注册用户
	if err := DB.Create(&req).Error; err != nil {
		utils.ReturnJSON(c, 500, "新增用户出错")
		return
	}
	utils.ReturnJSON(c, 0, "注册成功", gin.H{
		"user_id":  req.Id,
		"username": req.Username,
	})
}

// DeleteUser 删除用户DELETE
func DeleteUser(c *gin.Context) {
	var req models.User
	id := c.Param("id")

	if err := DB.First(&req, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ReturnJSON(c, 1, "未找到用户id")
			return
		}
	}
	if err := DB.Delete(&req, id).Error; err != nil {
		utils.ReturnJSON(c, 1, "删除用户失败")
		return
	}
	utils.ReturnJSON(c, 0, "删除成功")
}

// UpdateUser 更改用户数据put
func UpdateUser(c *gin.Context) {
	var req models.UpdateUser
	//检查是否有该id
	id := c.Param("id")
	if err := DB.First(&req, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ReturnJSON(c, 1, "未找到用户id")
			return
		}
	}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  utils.PareJSONError(err),
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
	user := models.User{}
	//将字符串转为int
	bigId, _ := strconv.Atoi(id)
	//将int转为int64
	req.Id = uint64(bigId)
	//检查用户名是否重复
	if err := DB.Model(models.User{}).Where("username = ?", req.Username).First(&req).Error; err == nil {
		fmt.Println(err)
		utils.ReturnJSON(c, 10086, "用户名已存在")
		return
	}
	//更新数据库
	if err := DB.Model(&user).Where("id = ?", req.Id).Updates(&req).Error; err != nil {
		fmt.Println(err)
		utils.ReturnJSON(c, 201, "修改失败")
		return
	}
	utils.ReturnJSON(c, 0, "修改成功")
}

// LoginUser 登录请求函数post
func LoginUser(c *gin.Context) {
	var req models.User
	// 1. 获取 JSON
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  utils.PareJSONError(err),
			"data": "参数错误",
		})
		return
	}
	// 2. 查询用户
	var user models.User
	if err := DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(401, gin.H{
				"code": 1,
				"msg":  "未找到该用户,请先注册",
				"data": nil,
			})
			return
		}
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
	token, _ := utils.GenerateToken(user.Id, user.Password)
	// 4. 成功
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func SelectUser(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		var user models.User
		if err := DB.
			Model(models.User{}).
			Select("id", "username", "avatar", "nickname", "createdAt", "updatedAt").
			Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(401, gin.H{
				"code": 1,
				"msg":  "未找到该用户",
			})
			return
		}
		utils.ReturnJSON(c, 0, "查询成功", user)
		return
	} else {
		var res []models.User
		if err := DB.
			Select("id", "username", "avatar", "nickname", "createdAt", "updatedAt").
			Find(&res).Error; err != nil {
			c.JSON(500, gin.H{
				"code": 1,
				"msg":  "查找用户失败",
			})
		}
		utils.ReturnJSON(c, 1, "查找成功", res)
		return
	}
}

// SelectUser 查询用户:输入查询的值进行查询
func SelectUsers(c *gin.Context) {
	id := c.Query("id")
	username := c.Query("username")
	if id != "" {
		var user models.User
		if err := DB.
			Select("id", "username", "avatar", "nickname", "createdAt", "updatedAt").
			Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(401, gin.H{
				"code": 1,
				"msg":  "查询用户失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "查询用户成功",
			"data": user,
		})
		return
	} else if username != "" {
		var res []models.User
		if err := DB.
			Select("id", "username", "avatar", "nickname", "createdAt", "updatedAt").
			Where("username LIKE ?", "%"+username+"%").Find(&res).Error; err != nil {
			c.JSON(500, gin.H{
				"code": 1,
				"msg":  "查找用户失败",
			})
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "查询成功",
			"data": res,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "请提供要查询的值",
		})
	}
}
