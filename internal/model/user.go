package model

import "time"

type User struct {
	Id int `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"type:varchar(255);not null;" json:"name"`
	LoginId   string `gorm:"type:varchar(255);not null;uniqueIndex:unique_login_id" json:"login_id"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}