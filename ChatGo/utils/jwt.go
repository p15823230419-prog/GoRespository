package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("7K7l3q9v8F2mNx5PbRtYs1HcDgJeLfKwUoQzEaBnWpXyGvRiSjTlOmNcCuZdXhA1kB6f9s2v4r8n5j3p0q7t1u")

type Claims struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(userID uint, username string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // token 24h 过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "chatgo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
