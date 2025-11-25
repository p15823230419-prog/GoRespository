package service

import (
	"abc/dto"
	"abc/entity"
	"abc/model"
)

// 注册请求体转为gorm实体
func RegisterRequestToEntity(request *dto.RegisterRequest, roles []model.Role) *entity.User {
	return &entity.User{
		Username: request.Username,
		Nickname: request.Nickname,
		Password: request.Password,
		Avatar:   request.Avatar,
		Phone:    request.Phone,
		Roles:    roles,
	}
}

// 添加请求体转换为实体
func CreateRoleRequestToEntity(request *dto.CreateRoleRequest) *model.Role {
	return &model.Role{
		RoleName: request.RoleName,
	}
}

// 实体转换为返回体
func EntityToSelectResponse(user entity.User) *dto.SelectResponse {
	resRoles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		resRoles[i] = role.RoleName
	}
	return &dto.SelectResponse{
		UserId:   user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
		Email:    user.Email,
		Roles:    resRoles,
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
