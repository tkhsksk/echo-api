package models

import (
	"time"
)

type Product struct {
    ID         uint      `gorm:"primaryKey"`
    Name       string    `gorm:"size:100;not null"`
    Price      float64   `gorm:"not null"`
    Content    *string   `gorm:"type:text"`
    Status     string    `gorm:"default:active"` // active / suspended など
    Images     []Image   `gorm:"many2many:product_images;"`

    CategoryID uint      // 外部キー
    Category   Category  // Categoryとのリレーション

    AdminID    uint      // 外部キー
    Admin      Admin     // Adminとのリレーション

    CreatedAt  time.Time
    UpdatedAt  time.Time
}