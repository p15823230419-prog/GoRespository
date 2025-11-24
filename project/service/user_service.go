package service

import (
	"abc/dao"
	"abc/dto"
	"abc/entity"
	"abc/utils"
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
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
	uid, err := u.userDao.Create(ctx, *RegisterRequestToEntity(req))
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
func (u *UserService) List(ctx *gin.Context) ([]*dto.SelectResponse, error) {
	username := ctx.Query("username")
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	pageSize, _ := strconv.Atoi(pageSizeStr)
	pageNum, _ := strconv.Atoi(pageNumStr)
	users, err := u.userDao.FindUsers(ctx, username, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return EntityToSelectResponses(users), nil
}

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
		Role:     user.Role,
	}
}

func EntityToSelectResponses(users []*entity.User) []*dto.SelectResponse {
	res := make([]*dto.SelectResponse, len(users))
	for i, user := range users {
		res[i] = EntityToSelectResponse(*user)
	}
	return res
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}
