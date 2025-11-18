package models

type Message struct {
	Id         string `gorm:"type:bigint(20);primary_key" json:"id"`
	SenderId   string `gorm:"type:bigint(20)" json:"sender"`
	ReceiverId string `gorm:"type:bigint(20)" json:"receiver"`
	Content    string `gorm:"type:varchar(511)" json:"content"`
	CreateAt   string `gorm:"" json:"create_at"`
	UpdateAt   string `gorm:"" json:"update_at"`
}
