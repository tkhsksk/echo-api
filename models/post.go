package models

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;foreignKey"`  // 外部キーの指定
	User      User      `gorm:"foreignKey:UserID"`    // リレーションの指定
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	Status    string    `gorm:"default:active"` // active / suspended など
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
