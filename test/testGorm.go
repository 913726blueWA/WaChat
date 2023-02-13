package main

import (
	"WaChat/models"
	"WaChat/utils"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	utils.InitConfig()
	fmt.Println(viper.GetString("mysql.dns"))
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Community{})
	//user := &models.UserBasic{}
	//user.Name = "Wa"
	//db.Create(user)
	db.First(&models.UserBasic{}, 1)
}
