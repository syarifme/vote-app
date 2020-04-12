package models

import "github.com/jinzhu/gorm"

type Answer struct {
	gorm.Model
	Title      string `json:"title"`
	QuestionID int    `json:"question_id"`
	IsAnswer   bool   `json:"is_answer"`
}
