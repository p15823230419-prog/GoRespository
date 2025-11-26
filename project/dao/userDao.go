package dao

import (
	"abc/entity"
	"abc/model"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// 注册用户
func (dao *UserDao) Create(ctx context.Context, user entity.User) (*uint64, error) {
	var modelUser = userEntityToModel(user)
	log.Println(modelUser)
	err := db.WithContext(ctx).Create(&modelUser).Error
	if err != nil {
		return nil, err
	}
	return &modelUser.Id, nil
}

// 通过id查询
func (dao *UserDao) FindById(ctx context.Context, id uint64) (*entity.User, error) {
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
func (dao *UserDao) FindByName(ctx context.Context, username string) (*entity.User, error) {
	var modelUser model.User
	err := db.WithContext(ctx).Where("username = ?", username).First(&modelUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return userModelToEntity(modelUser), nil
}

// 获取用户列表
func (dao *UserDao) FindAll(ctx context.Context, username string, pageNum int, pageSize int) ([]*entity.User, error) {
	var modelUsers []model.User

	query := db.WithContext(ctx).
		Model(&model.User{}).
		Preload("Roles")

	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}

	err := query.
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&modelUsers).Error

	if err != nil {
		return nil, err
	}

	return userModelsToEntities(modelUsers), nil
}

// 删除用户
func (dao *UserDao) Delete(ctx context.Context, id uint64) error {
	if err := db.WithContext(ctx).Delete(&model.User{}, "id", id).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 更改用户信息
func (dao *UserDao) Update(ctx context.Context, user entity.User, roleIds []uint64) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var roles []model.Role
		if len(roleIds) > 0 {
			if err := tx.Where("id IN ?", roleIds).Find(&roles).Error; err != nil {
				return err
			}
		}
		modelUser := userEntityToModel(user)

		// 1. 更新基础信息
		if err := tx.Updates(&modelUser).Error; err != nil {
			return err
		}

		// 2. 更新角色
		if err := tx.Model(&modelUser).Association("Roles").Replace(roles); err != nil {
			return err
		}

		return nil
	})
}
