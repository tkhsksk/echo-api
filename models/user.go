package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"` // ハッシュ化されたパスワード
	Status    string    `gorm:"default:active"` // active / suspended など
	CreatedAt time.Time
	UpdatedAt time.Time

	Posts        []Post        `gorm:"foreignKey:UserID"`
	UserSessions []UserSession `gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}