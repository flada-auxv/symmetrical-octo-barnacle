package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	UserID      int
	User        User
	Title       string
	Description string
	StartedAt   time.Time
	CompletedAt time.Time
}
