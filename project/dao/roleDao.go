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
func (rd *RoleDao) FindByRoleName(ctx context.Context, roleName string) (*model.Role, error) {
	var modelRole model.Role
	err := db.WithContext(ctx).Where("role_name = ?", roleName).First(&modelRole).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &modelRole, nil
}

func (rd RoleDao) CreateRole(ctx context.Context, req model.Role) error {
	if err := db.WithContext(ctx).Model(model.Role{}).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (rd RoleDao) UpdateRole(ctx context.Context, req model.Role) error {
	if err := db.WithContext(ctx).Model(model.Role{}).Save(&req).Error; err != nil {
		return err
	}
	return nil
}
