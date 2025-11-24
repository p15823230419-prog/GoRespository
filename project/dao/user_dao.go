package dao

import (
	"abc/entity"
	"abc/model"
	"abc/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserDao struct {
}

func (u *UserDao) Add(ctx context.Context, user entity.User) error {
	return utils.DB.WithContext(ctx).Create(userEntityToModel(user)).Error
}
func (u *UserDao) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var m model.User
	err := utils.DB.WithContext(ctx).Where("username = ?", username).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return userModelToEntity(m), nil
}
func userModelToEntity(m model.User) *entity.User {
	return &entity.User{
		Id:        m.Id,
		Username:  m.Username,
		Nickname:  m.Nickname,
		Password:  m.Password,
		Mobile:    m.Mobile,
		Email:     m.Email,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
	}
}
func userEntityToModel(e entity.User) *model.User {
	return &model.User{
		Id:        e.Id,
		Username:  e.Username,
		Nickname:  e.Nickname,
		Password:  e.Password,
		Mobile:    e.Mobile,
		Email:     e.Email,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
	}
}
func NewUserDao() *UserDao {
	return &UserDao{}
}
