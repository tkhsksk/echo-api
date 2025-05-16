package models

import (
	"time"
)

type APILog struct {
	ID        uint      `gorm:"primaryKey"`
	Method    string
	ClientIP  string
	Path      string
	Message   string 
	CreatedAt time.Time
}

type BlockLog struct {
	ID        uint      `gorm:"primaryKey"`
	ClientIP  string
	Path      string
	CreatedAt time.Time
}

