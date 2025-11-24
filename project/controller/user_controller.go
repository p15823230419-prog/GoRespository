package controller

import (
	"abc/dto"
	"abc/service"
	"abc/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

// 用户注册接口
func (u *UserController) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		utils.ReturnError(c, err)
		return
	}
	data, err := u.userService.Register(c, &registerRequest)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "注册成功", data)
}

// 用户登录接口
func (u *UserController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		utils.ReturnError(c, err)
		return
	}
	data, err := u.userService.Login(c, loginRequest)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, data)
}

// 查询接口
func (u *UserController) List(c *gin.Context) {
	data, err := u.userService.List(c)
	if err != nil {
		utils.ReturnError(c, err)
		return
	}
	utils.ReturnSuccess(c, "获取成功", gin.H{
		"list":  data,
		"total": len(data),
	})
}
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}
