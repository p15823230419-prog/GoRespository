package controller

import (
	"abc/dto"
	"abc/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func (u *UserController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	data, err := u.userService.Login(c, loginRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"data": data})
	}
}
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}
