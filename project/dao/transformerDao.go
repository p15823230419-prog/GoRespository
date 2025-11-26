package dao

import (
	"abc/dto"
	"abc/entity"
	"abc/model"
)

// 用户数据转实体
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
		Roles:     m.Roles,
	}
}

// 用户数据转实体
func userModelsToEntities(models []model.User) []*entity.User {
	res := make([]*entity.User, len(models))
	for i, m := range models {
		res[i] = userModelToEntity(m)
	}
	return res
}

// 用户实体转数据
func userEntityToModel(e entity.User) *model.User {
	roles := make([]model.Role, 0, len(e.Roles))
	for _, r := range e.Roles {
		roles = append(roles, model.Role{
			Id:       r.Id,
			RoleName: r.RoleName,
		})
	}
	return &model.User{
		Id:        e.Id,
		Username:  e.Username,
		Nickname:  e.Nickname,
		Password:  e.Password,
		Phone:     e.Phone,
		Email:     e.Email,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		Roles:     e.Roles,
	}
}

func creatMenuReqToModel(r *dto.CreateMenuReq) *model.Menu {
	return &model.Menu{
		Name:     r.Name,
		ParentId: r.ParentId,
	}
}
