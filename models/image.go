package models

import (
	"time"
)

type Image struct {
    ID       uint
    URL      string
    AltText  string
    IsPublic bool   `gorm:"default:true"` // 公開フラグ（true: 公開、false: 非公開）
    // 逆方向の関係
    Products      []Product      `gorm:"many2many:product_images;"`
    Notifications []Notification `gorm:"many2many:notification_images;"`

    CreatedAt time.Time
    UpdatedAt time.Time
}

