package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/syarifme/vote-app/models"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/vote_app?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connect")

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Answer{}, &models.Category{}, &models.Question{}, &models.Score{}, &models.User{}, &models.UserScore{})
}
