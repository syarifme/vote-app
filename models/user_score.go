package models

import "github.com/jinzhu/gorm"

type UserScore struct {
	gorm.Model
	UserID     int `json:"user_id"`
	TotalScore int `json:"total_score"`
}
