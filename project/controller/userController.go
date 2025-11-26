package controller

import (
	"abc/dto"
	"abc/service"
	"abc/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// 用户注册接口
func (uc *UserController) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		utils.ReturnBindError(c, err)
		return
	}
	data, err := uc.userService.Register(c, &registerRequest)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "注册成功", *data)
}

// 用户登录接口
func (uc *UserController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		utils.ReturnBindError(c, err)
		return
	}
	data, err := uc.userService.Login(c, loginRequest)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "登录成功", data)
}

// 退出登录接口
func (uc *UserController) Logout(c *gin.Context) {
	utils.ReturnSuccess(c, "退出成功")
}

// 查询接口
func (uc *UserController) List(c *gin.Context) {
	var selectRequest dto.SelectRequest
	if err := c.ShouldBind(&selectRequest); err != nil {
		utils.ReturnBindError(c, err)
	}
	data, err := uc.userService.List(c, selectRequest)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "获取成功", gin.H{
		"list":     data,
		"total":    len(data),
		"pageNum":  c.Query("pageNum"),
		"pageSize": c.Query("pageSize"),
	})
}

// 删除用户接口
func (uc *UserController) Delete(c *gin.Context) {
	if err := uc.userService.Delete(c); err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "删除成功")
}

// 更新用户信息接口
func (uc *UserController) Update(c *gin.Context) {
	var updateRequest dto.UpdateRequest
	if err := c.ShouldBind(&updateRequest); err != nil {
		fmt.Println(err)
		utils.ReturnBindError(c, err)
		return
	}
	if err := uc.userService.Update(c, updateRequest); err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "更新成功")
}
