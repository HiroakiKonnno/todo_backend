package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Task struct {
	ID        int      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:varchar(255)" json:"content"`
	UserId    int       `gorm:"type:integer;not null" json:"user_id"`
	StartDate *time.Time `gorm:"type:date" json:"start_date"`
	EndDate *time.Time `gorm:"type:date" json:"end_date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t Task) Validate() error {
	return validation.ValidateStruct(&t, 
		validation.Field(&t.Title, validation.Required),
	)
}