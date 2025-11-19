package models

import "time"

type Message struct {
	Id         uint64    `gorm:"type:bigint(20);primary_key" json:"id"`
	SenderId   uint64    `gorm:"type:bigint(20)" json:"senderId"`
	ReceiverId uint64    `gorm:"type:bigint(20)" json:"receiverId"`
	Content    string    `gorm:"type:varchar(511)" json:"content"`
	CreatedAt  time.Time `gorm:"type:datetime;comment:创建时间;" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"type:datetime;comment:更新时间;" json:"updatedAt"`
}
