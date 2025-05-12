package models

import (
	"time"
)

type Product struct {
    ID            uint      `gorm:"primaryKey"`
    Name          string    `gorm:"size:100;not null"`
    Price         float64   `gorm:"not null"`
    Content       *string   `gorm:"type:text"`
    Status        string    `gorm:"default:active"` // active / suspended など

    CategoryID    uint      // 外部キー
    Category      Category  // Categoryとのリレーション

    AdminID    uint      // 外部キー
    Admin      Admin     // Adminとのリレーション

    CreatedAt     time.Time
    UpdatedAt     time.Time
}

type ProductForUser struct {
    ID      uint    `json:"id"`
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    Content string  `json:"content"`
    CategoryID uint `json:"categoryID"`
}