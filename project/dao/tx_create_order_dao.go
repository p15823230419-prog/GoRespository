package dao

import (
	"abc/utils"
	"context"
	"gorm.io/gorm"
)

type TxCreateOrderDao struct {
}

func (t *TxCreateOrderDao) CreateOrder(ctx context.Context) error {
	// 不建议使用这种事务方式
	//utils.DB.Begin()
	//utils.DB.Commit()
	//utils.DB.Rollback()
	err := utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return nil
	})
	return err
}
