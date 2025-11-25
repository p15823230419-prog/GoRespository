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
	utils.ReturnSuccess(c, "注册成功", *data)
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
	utils.ReturnSuccess(c, "登录成功", data)
}

// 退出登录接口
func (u *UserController) Logout(c *gin.Context) {
	utils.ReturnSuccess(c, "退出成功")
}

// 查询接口
func (u *UserController) List(c *gin.Context) {
	data, err := u.userService.List(c)
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
func (u *UserController) Delete(c *gin.Context) {
	if err := u.userService.Delete(c); err != nil {
		utils.ReturnError(c, err)
	}
	utils.ReturnSuccess(c, "删除成功")
}
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// 更新用户信息接口
func (u *UserController) Update(c *gin.Context) {
	var updateRequest dto.UpdateRequest
	if err := c.ShouldBind(&updateRequest); err != nil {
		utils.ReturnError(c, err)
	}
	if err := u.userService.Uptate(c, updateRequest); err != nil {

	}
	utils.ReturnSuccess(c, "更新成功")
}
