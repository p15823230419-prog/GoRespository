package dao

import (
	"abc/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type RoleDao struct {
}

func NewRoleDao() *RoleDao {
	return &RoleDao{}
}

// 查询
func (dao *RoleDao) FindByName(ctx context.Context, roleName string) (*model.Role, error) {
	var modelRole model.Role
	err := db.WithContext(ctx).Where("role_name = ?", roleName).First(&modelRole).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &modelRole, nil
}

func (dao *RoleDao) FindAll(ctx context.Context) ([]model.Role, error) {
	var modelRole []model.Role
	err := db.WithContext(ctx).Find(&modelRole).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return modelRole, nil
}

func (dao *RoleDao) Create(ctx context.Context, req model.Role) error {
	if err := db.WithContext(ctx).Model(model.Role{}).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (dao *RoleDao) Update(ctx context.Context, req model.Role) error {
	if err := db.WithContext(ctx).Model(model.Role{}).Save(&req).Error; err != nil {
		return err
	}
	return nil
}
