package responses

import (
	"time"
)

type UserSession struct {
    ID        string    `json:"id"`
    UserID    uint      `json:"userId"`
    CreatedAt time.Time `json:"createdAt"`
    ExpiresAt time.Time `json:"expiresAt"`
}

type UserSessionByUserID struct {
    ID        string    `json:"id"`
    CreatedAt time.Time `json:"createdAt"`
    ExpiresAt time.Time `json:"expiresAt"`
}