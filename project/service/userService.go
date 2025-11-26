package service

import (
	"abc/dao"
	"abc/dto"
	"abc/model"
	"abc/utils"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}

// 注册功能实现
func (service *UserService) Register(c *gin.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	ctx := c.Request.Context()
	// 查找用户名是否重复
	if user, err := service.userDao.FindByName(ctx, req.Username); err != nil {
		return nil, err
	} else if user != nil {
		return nil, errors.New("用户名已存在")
	}
	log.Println(req)
	// 添加角色
	roles := make([]model.Role, len(req.RoleIDs))
	for i, id := range req.RoleIDs {
		roles[i] = model.Role{Id: id}
	}
	log.Println(roles, req.RoleIDs)
	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}
	req.Password = hashedPassword
	// 创建账户
	uid, err := service.userDao.Create(ctx, *RegisterRequestToEntity(req, roles))
	if err != nil {
		return nil, err
	}

	res := &dto.RegisterResponse{
		UserId: *uid,
	}
	return res, nil
}

// 登录功能实现
func (service *UserService) Login(c *gin.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1.查询用户
	user, err := service.userDao.FindByName(c.Request.Context(), req.Username)
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
func (service *UserService) List(c *gin.Context, req dto.SelectRequest) ([]*dto.SelectResponse, error) {

	users, err := service.userDao.FindAll(c.Request.Context(), req.Username, req.PageNum, req.PageSize)
	if err != nil {
		log.Println(err)
		return nil, errors.New("数据库错误")
	}

	return EntityToSelectResponses(users), nil
}

// 删除用户
func (service *UserService) Delete(c *gin.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("请输入要删除的id")
	}
	userId, _ := strconv.Atoi(id)

	if user, err := service.userDao.FindById(c.Request.Context(), uint64(userId)); err != nil {
		return errors.New("数据库错误")
	} else if user == nil {
		return errors.New("未找到用户")
	}

	if err := service.userDao.Delete(c, uint64(userId)); err != nil {
		return errors.New("删除失败,数据库错误")
	}
	return nil
}

// 更新用户信息
func (service *UserService) Update(c *gin.Context, req dto.UpdateRequest) error {
	// 查找id
	user, err := service.userDao.FindById(c.Request.Context(), req.Id)
	if err != nil {
		return errors.New("查找id错误")
	} else if user == nil {
		return errors.New("未找到用户,请先注册")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	req.Password = hashedPassword

	// 更新数据
	if err = service.userDao.Update(c.Request.Context(), *UpdateRequestToEntity(req), req.RoleIDs); err != nil {
		fmt.Println(err)
		return errors.New("更新数据库错误")
	}
	return nil
}
