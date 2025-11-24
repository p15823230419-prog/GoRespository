package middleware

import (
	"ChatGo/controllers"
	"ChatGo/models"
	"ChatGo/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ReturnJSON(c, 401, "token为空")
			c.Abort()
			return
		}

		// 解析 token
		claims, err := utils.ParseToken(authHeader)
		if err != nil {
			var ve *jwt.ValidationError

			if errors.As(err, &ve) {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					utils.ReturnJSON(c, 401, "解析失败")
				}
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					utils.ReturnJSON(c, 401, "登录已过期请重新登录")
				}
				if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
					c.JSON(401, "token 签名错误")
				}
				c.Abort()
				return
			}
			utils.ReturnJSON(c, 401, "token解析错误")
			c.Abort()
			return
		}
		//判断用户信息有效性
		user := models.User{}
		if err = controllers.DB.First(&user, "password = ? AND id = ?", claims.Password, claims.UserID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utils.ReturnJSON(c, 1, "用户信息错误,请重新登录")
				c.Abort()
				return
			}
		}
		// 把用户信息放进 context
		c.Set("userId", claims.UserID)
		c.Set("username", claims.Password)

		c.Next()
	}
}
