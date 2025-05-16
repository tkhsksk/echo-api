package utils

import (
    "log"
    "github.com/labstack/echo/v4"
)

var ANSIColors = map[int]string{
    001: "\033[0m",
    002: "\033[31m",// Caution
    003: "\033[38;5;072m",// Success
    004: "\033[38;5;185m",// Warn
    005: "\033[38;5;069m",// Blue
}

func LogRequest(c echo.Context, level int, msg string) {
    color := ANSIColors[level]
    reset := ANSIColors[001]
    method := c.Request().Method
    path := c.Path()

    log.Println(color + "[" + method + " -> " + path + "]:" + reset + msg)
}
