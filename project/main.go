package main

import (
	"abc/controller"
	"abc/dao"
	"abc/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

var userController *controller.UserController

func WebInit() {
	//初始化
	userController = controller.NewUserController()
	r := gin.Default()
	// 用户登录
	r.POST("/user/login", userController.Login)
	// 注册用户
	r.POST("/user/register", userController.Register)
	//创建user用户组
	userGroup := r.Group("/user", middleware.JWTAuth())

	// 查询用户
	userGroup.GET("/list", userController.List)
	// 退出登录
	userGroup.GET("/logout", userController.Logout)
	// 删除用户
	userGroup.DELETE(":id", userController.Delete)
	//监听8080端口
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	dao.InitDB()
	WebInit()
}
