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

type UserSessionResponse struct {
	ID        string 	`json:"id"`
	UserID    uint   	`json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type UserSessionResponseByUserID struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type AdminSession struct {
	ID        string    `gorm:"primaryKey"` // UUIDとかで生成
	AdminID   uint      // 外部キー
    Admin     Admin     // Adminとのリレーション
	ExpiresAt time.Time
	CreatedAt time.Time
}
