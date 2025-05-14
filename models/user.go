package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"` // ハッシュ化されたパスワード
	Status    string    `gorm:"default:active"` // active / suspended など
	CreatedAt time.Time
	UpdatedAt time.Time

	Posts        []Post        `gorm:"foreignKey:UserID"`
	UserSessions []UserSession `gorm:"foreignKey:UserID"`
}