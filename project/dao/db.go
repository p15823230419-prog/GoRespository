package dao

import (
	"abc/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止自动复数
		},
	})
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	log.Printf("连接成功")
	db = gormDb
	err = db.AutoMigrate(&model.User{}, &model.Role{}, &model.Menu{})
	if err != nil {
		log.Fatal(err)
		return
	}
}
