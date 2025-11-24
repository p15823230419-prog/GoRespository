package dao

import (
	"abc/entity"
	"abc/model"
	"abc/utils"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserDao struct {
}

// 注册用户
func (u *UserDao) Create(ctx context.Context, user entity.User) (*uint64, error) {
	var modelUser = userEntityToModel(user)
	err := utils.DB.WithContext(ctx).Create(&modelUser).Error
	if err != nil {
		return nil, err
	}
	return &modelUser.Id, nil
}

// 通过名字查询
func (u *UserDao) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var modelUser model.User
	err := utils.DB.WithContext(ctx).Where("username = ?", username).First(&modelUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return userModelToEntity(modelUser), nil
}

// 获取用户列表
func (u *UserDao) FindUsers(ctx *gin.Context, username string, pageNum int, pageSize int) ([]*entity.User, error) {
	if username != "" {
		var modelUsers []model.User
		if err := utils.DB.WithContext(ctx).
			Model(model.User{}).
			Select("id", "username", "avatar", "email", "phone", "role", "nickname", "createdAt", "updatedAt").
			Where("username like ?", "%"+username+"%").
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).
			Find(&modelUsers).
			Error; err != nil {
			return nil, err
		}
		return userModelsToEntities(modelUsers), nil
	} else {
		var modelUsers []model.User
		if err := utils.DB.WithContext(ctx).
			Model(model.User{}).
			Select("id", "username", "avatar", "email", "phone", "role", "nickname", "createdAt", "updatedAt").
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).
			Find(&modelUsers).
			Error; err != nil {
			return nil, err
		}
		return userModelsToEntities(modelUsers), nil
	}

}

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

func userEntityToModel(e entity.User) *entity.User {
	return &entity.User{
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
func NewUserDao() *UserDao {
	return &UserDao{}
}
