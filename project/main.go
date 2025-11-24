package main

import (
	"abc/controller"
	"abc/middleware"
	"abc/utils"
	"log"

	"github.com/gin-gonic/gin"
)

var userController *controller.UserController

func WebInit() {
	//初始化
	userController = controller.NewUserController()
	r := gin.Default()
	userGroup := r.Group("/user", middleware.JWTAuth())
	r.POST("/user/login", userController.Login)
	r.POST("/user/register", userController.Register)

	userGroup.GET("/list", userController.List)

	//监听8080端口
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	utils.InitDB()
	WebInit()
}
