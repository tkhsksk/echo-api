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

