package models

import "github.com/jinzhu/gorm"

type Score struct {
	gorm.Model
	UserID     int `json:"user_id"`
	QuestionID int `json:"question_id"`
	AnswerID   int `json:"answer_id"`
}
