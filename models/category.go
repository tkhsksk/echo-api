package models

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
    ID        uint           `gorm:"primaryKey"`
    Name      string         `gorm:"size:100;not null"`
    Content   *string        `gorm:"type:text"`
    Status    string         `gorm:"default:active"` // active / suspended など
    AdminID   uint           // 外部キー
    Admin     Admin          // Adminとのリレーション
    // 階層構造のための自己参照リレーション
    ParentID  *uint          // null を許容
    Parent    *Category      `gorm:"foreignKey:ParentID"`
    Children  []Category     `gorm:"foreignKey:ParentID"`

    Products  []Product      `gorm:"foreignKey:CategoryID"`
    
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}