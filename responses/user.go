package responses

import (
	"time"
)

type User struct {
    ID        uint   `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Status    string `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}