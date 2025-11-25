package service

import (
	"abc/dao"
	"abc/dto"
	"abc/model"
	"abc/utils"
	"errors"
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
func (u *UserService) Register(c *gin.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	ctx := c.Request.Context()
	// 查找用户名是否重复
	if user, err := u.userDao.FindByUsername(ctx, req.Username); err != nil {
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
	uid, err := u.userDao.Create(ctx, *RegisterRequestToEntity(req, roles))
	if err != nil {
		return nil, err
	}

	res := &dto.RegisterResponse{
		UserId: *uid,
	}
	return res, nil
}

// 登录功能实现
func (u *UserService) Login(c *gin.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1.查询用户
	user, err := u.userDao.FindByUsername(c.Request.Context(), req.Username)
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
	users, err := u.userDao.FindUsers(ctx.Request.Context(), username, pageNum, pageSize)
	if err != nil {
		log.Println(err)
		return nil, errors.New("数据库错误")
	}
	return EntityToSelectResponses(users), nil
}

// 删除用户
func (u *UserService) Delete(c *gin.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("请输入要删除的id")
	}
	userId, _ := strconv.Atoi(id)
	if user, err := u.userDao.FindById(c.Request.Context(), uint64(userId)); err != nil {
		return errors.New("数据库错误")
	} else if user == nil {
		return errors.New("未找到用户,请先注册")
	}
	if err := u.userDao.DeleteUser; err != nil {
		return errors.New("删除失败,数据库错误")
	}
	return nil
}

// 更新用户信息
func (u *UserService) Uptate(c *gin.Context, req dto.UpdateRequest) error {
	user, err := u.userDao.FindById(c.Request.Context(), req.Id)
	if err != nil {
		return err
	} else if user == nil {
		return errors.New("未找到用户,请先注册")
	}
	if err = u.userDao.Update(c.Request.Context(), req); err != nil {
		return err
	}
	return nil
}
