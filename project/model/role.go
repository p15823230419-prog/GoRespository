package model

import "time"

type Role struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	RoleName  string    `gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time `gorm:"type:datetime;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime;comment:更新时间" json:"updatedAt"`
}
