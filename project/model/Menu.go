package model

import "time"

type Menu struct {
	Id        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	ParentId  uint64    `json:"parentId" gorm:"default:0"`
	Name      string    `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"createdAt"`
	Children  []Menu    `json:"children" gorm:"-"`
}
