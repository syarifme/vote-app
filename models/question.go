package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	CategoryID int        `json:"category_id"`
	Title      string     `json:"title"`
	Score      int        `json:"score"`
	StartDate  *time.Time `json:"start_date"`
	EndDate    *time.Time `json:"end_date"`
	Status     string     `json:"status"`
}
