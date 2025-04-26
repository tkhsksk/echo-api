package models

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	User      User
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
