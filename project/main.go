package main

import (
	"abc/controller"

	"github.com/gin-gonic/gin"
)

var userController *controller.UserController

func main() {
	r := gin.Default()
	userController = controller.NewUserController()
	// gin 框架 官方中间件
	r.GET("/user/login", userController.Login)
	//监听端口默认为8080
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
