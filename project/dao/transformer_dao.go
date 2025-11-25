package dao

import (
	"abc/entity"
	"abc/model"
)

// 模型转换
func userModelToEntity(m model.User) *entity.User {
	return &entity.User{
		Id:        m.Id,
		Username:  m.Username,
		Nickname:  m.Nickname,
		Password:  m.Password,
		Phone:     m.Phone,
		Email:     m.Email,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
	}
}

func userModelsToEntities(models []model.User) []*entity.User {
	res := make([]*entity.User, len(models))
	for i, m := range models {
		res[i] = userModelToEntity(m)
	}
	return res
}

func userEntityToModel(e entity.User) *model.User {
	return &model.User{
		Id:        e.Id,
		Username:  e.Username,
		Nickname:  e.Nickname,
		Password:  e.Password,
		Phone:     e.Phone,
		Email:     e.Email,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
	}
}
