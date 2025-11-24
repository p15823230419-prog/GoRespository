package service

import (
	"abc/dao"
	"abc/dto"
	"abc/entity"
	"abc/utils"
	"context"
	"errors"
)

type UserService struct {
	userDao *dao.UserDao
}

// 注册功能实现
func (u *UserService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {

	// 查找用户名是否重复
	if user, err := u.userDao.FindByUsername(ctx, req.Username); err != nil {
		return nil, err
	} else if user != nil {
		return nil, errors.New("用户名已存在")
	}
	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}
	req.Password = hashedPassword
	// 创建账户
	uid, err := u.userDao.Create(ctx, RegisterRequestToEntity(req))
	if err != nil {
		return nil, err
	}

	res := &dto.RegisterResponse{
		UserId: *uid,
	}
	return res, nil
}

func (u *UserService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1.查询用户
	user, err := u.userDao.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.New("未找到用户,请先注册")
	}
	// 2.验证账号密码
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("密码错误")
	}
	// 3.成功登录返回token
	token, _ := utils.GenerateToken(user.Id, user.Password)
	res := &dto.LoginResponse{Token: token}
	return res, nil
}

// 查询用户表
func (u *UserService) List(ctx context.Context) ([]*entity.User, error) {
	users, err := u.userDao.FindUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 注册请求体转为gorm实体
func RegisterRequestToEntity(response *dto.RegisterRequest) entity.User {
	return entity.User{
		Username: response.Username,
		Nickname: response.Nickname,
		Password: response.Password,
		Avatar:   response.Avatar,
		Phone:    response.Phone,
	}
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}
