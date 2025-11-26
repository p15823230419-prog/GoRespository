package dao

import (
	"abc/dto"
	"abc/model"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type MenuDao struct {
}

func NewMenuDao() *MenuDao {
	return &MenuDao{}
}

func (dao *MenuDao) FindById(ctx context.Context, id uint64) (*model.Menu, error) {
	var menu model.Menu

	err := db.WithContext(ctx).First(&menu, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &menu, nil
}

func (dao *MenuDao) FindByName(ctx context.Context, name string) (*model.Menu, error) {
	var modelMenu model.Menu

	err := db.WithContext(ctx).Where("name = ?", name).First(&modelMenu).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &modelMenu, nil
}

func (dao *MenuDao) FindAll(ctx context.Context) ([]model.Menu, error) {
	var modelMenu []model.Menu
	err := db.WithContext(ctx).Find(&modelMenu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return modelMenu, nil
}

func (dao *MenuDao) Create(ctx context.Context, req *dto.CreateMenuReq) error {
	return db.WithContext(ctx).Model(model.Menu{}).Create(creatMenuReqToModel(req)).Error
}
func (dao *MenuDao) Delete(ctx context.Context, id uint64) error {
	if err := db.WithContext(ctx).Delete(&model.Menu{}, "id", id).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
