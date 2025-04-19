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
