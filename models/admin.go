package models

import (
	"time"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"` // ハッシュ化されたパスワード
	Status    string    `gorm:"default:active"` // active / suspended など
	CreatedAt time.Time
	UpdatedAt time.Time

	AdminSessions []AdminSession `gorm:"foreignKey:AdminID"`
	Passcodes     []Passcode     `gorm:"foreignKey:AdminID"`
	Notifications []Notification `gorm:"foreignKey:AdminID"`
	Categories    []Category     `gorm:"foreignKey:AdminID"`
	Products      []Product      `gorm:"foreignKey:AdminID"`
}