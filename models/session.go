package models

import (
	"time"
)

type Session struct {
	ID        string    `gorm:"primaryKey"` // UUIDとかで生成
	UserID    uint      `gorm:"not null"`
	User      User
	ExpiresAt time.Time
	CreatedAt time.Time
}

type SessionResponse struct {
	ID        string `json:"id"`
	UserID    uint   `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type SessionUserResponse struct {
	ID        string `json:"id"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}
