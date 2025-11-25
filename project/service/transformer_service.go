package service

import (
	"abc/dto"
	"abc/entity"
)

// 注册请求体转为gorm实体
func RegisterRequestToEntity(request *dto.RegisterRequest) *entity.User {
	return &entity.User{
		Username: request.Username,
		Nickname: request.Nickname,
		Password: request.Password,
		Avatar:   request.Avatar,
		Phone:    request.Phone,
	}
}

// 实体转换为返回体
func EntityToSelectResponse(user entity.User) *dto.SelectResponse {
	return &dto.SelectResponse{
		UserId:   user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
		Email:    user.Email,
	}
}

// 实体数组转为返回体数组
func EntityToSelectResponses(users []*entity.User) []*dto.SelectResponse {
	res := make([]*dto.SelectResponse, len(users))
	for i, user := range users {
		res[i] = EntityToSelectResponse(*user)
	}
	return res
}
