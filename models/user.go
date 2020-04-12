package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name 		string	`json:"name"`
	Username 	string	`json:"username"	gorm:"type:varchar(50);unique"`
	Email		string	`json:"email"		gorm:"type:varchar(100);unique_index"`
	Password	string	`json:"password"`
}