package model

type Role struct {
	Id       uint64 `gorm:"primaryKey;autoIncrement"`
	RoleName string `gorm:"type:varchar(100);unique;not null"`
}
