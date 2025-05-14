package responses

import (
	"time"
)

type Category struct {
    ID         uint    `json:"id"`
    Name       string  `json:"name"`
    Status     string  `json:"status"`
    ParentID   *uint    `json:"parent_id"`
    Admin      AdminSummary `json:"admin"`
    CreatedAt  time.Time  `json:"created_at"`
    UpdatedAt  time.Time  `json:"updated_at"`
}

// カテゴリー情報取得
type CategorySummary struct {
    ID     uint   `json:"id"`
    Name   string `json:"name"`
}

type CategoryTree struct {
    ID       uint                   `json:"id"`
    Name     string                 `json:"name"`
    Status   string                 `json:"status"`
    Children []CategoryTree `json:"children"`
}