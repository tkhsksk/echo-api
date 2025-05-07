package models

import (
	"time"
)

type UserSession struct {
	ID        string    `gorm:"primaryKey"` // UUIDとかで生成
	UserID    uint      `gorm:"not null;foreignKey"`  // 外部キーの指定
	User      User      `gorm:"foreignKey:UserID"`    // リレーションの指定
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
	AdminID   uint      `gorm:"not null;foreignKey"`  // 外部キーの指定
	Admin     Admin     `gorm:"foreignKey:AdminID"`   // リレーションの指定
	ExpiresAt time.Time
	CreatedAt time.Time
}
