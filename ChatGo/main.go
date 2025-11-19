package main

import (
	"ChatGo/middleware"

	"ChatGo/controllers"
	"ChatGo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitValidator()
	//初始化数据库
	controllers.InitDB()
	//初始化gin
	r := gin.Default()
	//注册用户
	r.POST("/register", controllers.RegisterUser)
	//登录用户
	r.POST("/login", controllers.LoginUser)
	//中间件
	userGroup := r.Group("/user", middleware.JWTAuth())
	{
		//删除用户
		userGroup.DELETE("/delete", controllers.DeleteUser)
		//更改用户
		userGroup.PUT("/update", controllers.UpdateUser)
		//查询用户
		userGroup.GET("/select", controllers.SelectUser)
		//发送消息
		userGroup.POST("/send", controllers.SendHandler)
		//查看消息
		userGroup.GET("/messages", controllers.GetMessages)
	}
	//开始监听端口
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
