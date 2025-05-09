package models

import (
	"time"
)

type Passcode struct {
	ID        uint      `gorm:"primaryKey"`
	AdminID   uint      // 外部キー
    Admin     Admin     // Adminとのリレーション
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
