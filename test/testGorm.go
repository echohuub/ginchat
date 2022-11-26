package main

import (
	"github.com/heqingbao/ginchat/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("ginchat.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Message{})

	//user := &models.UserBasic{}
	//user.Name = "张三"
	//db.Create(user)
	//
	//fmt.Println(db.First(user, 1))
	//
	//db.Model(user).Update("Password", "1234")
}
