package models

import (
	"time"
)

type Passcode struct {
	ID        uint      `gorm:"primaryKey"`
	AdminID   uint      `gorm:"not null;foreignKey"`  // 外部キーの指定
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
