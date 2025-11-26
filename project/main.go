package main

import (
	"abc/controller"
	"abc/dao"
	"abc/middleware"
	"abc/utils"
	"log"

	"github.com/gin-gonic/gin"
)

var userController *controller.UserController
var roleController *controller.RoleController
var menuController *controller.MenuController

func WebInit() {
	//初始化
	userController = controller.NewUserController()
	roleController = controller.NewRoleController()
	menuController = controller.NewMenuController()
	r := gin.Default()
	// 用户登录
	r.POST("/user/login", userController.Login)
	// 注册用户
	r.POST("/user/register", userController.Register)
	// 创建user用户组
	userGroup := r.Group("/user", middleware.JWTAuth())
	// 查询用户
	userGroup.GET("/list", userController.List)
	// 退出登录
	userGroup.GET("/logout", userController.Logout)
	// 删除用户
	userGroup.DELETE(":id", userController.Delete)
	// 更新用户
	userGroup.PUT("", userController.Update)
	// 创建角色组
	roleGroup := r.Group("/role")
	// 添加角色
	roleGroup.POST("/add", roleController.Create)
	// 创建菜单组
	menuGroup := r.Group("/menu")
	menuGroup.POST("/add", menuController.Create)
	menuGroup.DELETE(":id", menuController.Delete)
	menuGroup.GET("list", menuController.List)
	//监听8080端口
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	utils.InitValidator()
	dao.InitDB()
	WebInit()
}
