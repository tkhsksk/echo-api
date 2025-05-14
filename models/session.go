package models

import (
	"time"
)

type UserSession struct {
	ID        string    `gorm:"primaryKey"` // UUIDとかで生成
	UserID    uint      // 外部キー
    User      User     // Userとのリレーション
	ExpiresAt time.Time
	CreatedAt time.Time
}

type AdminSession struct {
	ID        string    `gorm:"primaryKey"` // UUIDとかで生成
	AdminID   uint      // 外部キー
    Admin     Admin     // Adminとのリレーション
	ExpiresAt time.Time
	CreatedAt time.Time
}