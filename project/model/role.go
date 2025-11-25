package model

type Role struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	RoleName string
}
