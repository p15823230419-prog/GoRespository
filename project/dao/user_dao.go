package dao

import (
	"abc/dto"
	"abc/entity"
	"abc/model"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// 注册用户
func (u *UserDao) Create(ctx context.Context, user entity.User) (*uint64, error) {
	var modelUser = userEntityToModel(user)
	err := db.WithContext(ctx).Create(&modelUser).Error
	if err != nil {
		return nil, err
	}
	return &modelUser.Id, nil
}

// 通过id查询
func (u *UserDao) FindById(ctx context.Context, id uint64) (*entity.User, error) {
	var modelUser model.User
	err := db.WithContext(ctx).First(&modelUser, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return userModelToEntity(modelUser), nil
}

// 通过名字查询
func (u *UserDao) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var modelUser model.User
	err := db.WithContext(ctx).Where("username = ?", username).First(&modelUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return userModelToEntity(modelUser), nil
}

// 获取用户列表
func (u *UserDao) FindUsers(ctx context.Context, username string, pageNum int, pageSize int) ([]*entity.User, error) {
	if username != "" {
		var modelUsers []model.User
		if err := db.WithContext(ctx).
			Model(model.User{}).
			Select("id", "username", "avatar", "email", "phone", "nickname", "createdAt", "updatedAt").
			Where("username like ?", "%"+username+"%").
			Preload("Roles").
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).
			Find(&modelUsers).
			Error; err != nil {
			return nil, err
		}
		return userModelsToEntities(modelUsers), nil
	} else {
		var modelUsers []model.User
		if err := db.WithContext(ctx).
			Model(model.User{}).
			Select("id", "username", "avatar", "email", "phone", "nickname", "createdAt", "updatedAt").
			Preload("Roles").
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).
			Find(&modelUsers).
			Error; err != nil {
			return nil, err
		}
		return userModelsToEntities(modelUsers), nil
	}
}

// 删除用户
func (u *UserDao) DeleteUser(ctx *gin.Context) error {
	if err := db.WithContext(ctx).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

// 更改部分用户信息patch
func (u *UserDao) Update(ctx context.Context, user dto.UpdateRequest) error {
	if err := db.WithContext(ctx).
		Where("id = ?", user.Id).
		Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// 全量更新put
func (u *UserDao) Updates(ctx context.Context, user dto.UpdateRequest) error {
	if err := db.WithContext(ctx).
		Where("id = ?", user.Id).
		Save(user).Error; err != nil {
		return err
	}
	return nil
}
