package models

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      // 外部キー
    User      User      // Userとのリレーション
	Title     string    `gorm:"size:100;not null"`
	Content   string    `gorm:"type:text"`
	Status    string    `gorm:"default:active"` // active / suspended など
	CreatedAt time.Time
	UpdatedAt time.Time
}
