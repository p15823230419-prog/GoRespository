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
		userGroup.DELETE("/:id", controllers.DeleteUser)
		//全部更新用户
		userGroup.PUT("/:id", controllers.UpdateUser)
		//部分更新
		userGroup.PATCH("/:id", controllers.UpdateUser)
		//查询所有用户
		userGroup.GET("", controllers.SelectUser)
		//查询单个用户
		userGroup.GET("/:id", controllers.SelectUser)
		//发送消息
		userGroup.POST("/messages/:id", controllers.SendHandler)
		//查看消息
		userGroup.GET("/messages/:id", controllers.GetMessages)
	}
	//开始监听端口
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
