package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// クエリパラメータ "limit" を取得し、整数として返す（デフォルト値と上限を引数で指定）
func ParseLimitParam(c echo.Context, defaultLimit, maxLimit int) int {
	limit := defaultLimit
	if l := c.QueryParam("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil {
			limit = parsed
			if limit > maxLimit {
				limit = maxLimit
			}
		}
	}
	return limit
}


