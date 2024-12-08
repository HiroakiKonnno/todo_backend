package model

import "time"

type Task struct {
	ID          int       `gorm:"primaryKey"`
	Title       string    `gorm:"size:255"`
	Content     string    `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
