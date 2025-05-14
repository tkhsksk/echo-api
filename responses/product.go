package responses

import (
	"time"
    "api/models"
)

type ProductForUser struct {
    ID         uint    `json:"id"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    Content    string  `json:"content"`
    CategoryID uint    `json:"category_id"`
}

type ProductForAdmin struct {
    ID         uint    `json:"id"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    Status     string  `json:"status"`
    Category   CategorySummary `json:"category"`
    Admin      AdminSummary    `json:"admin"`
    CreatedAt  time.Time  `json:"created_at"`
    UpdatedAt  time.Time  `json:"updated_at"`
}

// api用の構造体
func NewProductForAdmin(p models.Product) ProductForAdmin {
    return ProductForAdmin{
        ID:     p.ID,
        Name:   p.Name,
        Price:  p.Price,
        Status: p.Status,
        Category: CategorySummary{
            ID:   p.Category.ID,
            Name: p.Category.Name,
        },
        Admin: AdminSummary{
            ID:     p.Admin.ID,
            Name:   p.Admin.Name,
            Status: p.Admin.Status,
        },
        CreatedAt: p.CreatedAt,
        UpdatedAt: p.UpdatedAt,
    }
}
