package responses

import (
	"time"
)

type Admin struct {
    ID        uint   `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Status    string `json:"status"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// 汎用的に admin の一部を含む構造体
type AdminSummary struct {
    ID     uint   `json:"id"`
    Name   string `json:"name"`
    Status string `json:"status"`
}
