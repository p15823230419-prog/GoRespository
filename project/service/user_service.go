package service

import (
	"abc/dao"
	"abc/dto"
	"abc/entity"
	"context"
)

type UserService struct {
	userDao *dao.UserDao
}

// dto 是什么
// dto vo po
//
//	context.Context
func (u *UserService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	return nil, nil
}
func (u *UserService) Add(ctx context.Context, req dto.AddRequest) error {
	_ = u.userDao.Add(ctx, entity.User{})
	return nil
}
func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}
