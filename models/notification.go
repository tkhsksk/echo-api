package models

import (
	"time"
	"gorm.io/gorm"
)

type Notification struct {
    ID        uint           `gorm:"primaryKey"`
    Title     string         `gorm:"not null"`
	Content   string         `gorm:"type:text"`
	Status    string         `gorm:"default:active"` // active / suspended など

    AdminID   uint           // 外部キー
    Admin     Admin          // Adminとのリレーション

    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
