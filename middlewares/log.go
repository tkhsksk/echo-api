package middlewares

import (
	"bytes"
	"encoding/json"
	"log"

	"net/http"
	"api/db"
	"api/models"
	"api/utils"

	"github.com/labstack/echo/v4"
)

type responseCaptureWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (w *responseCaptureWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // レスポンス内容をバッファに保存
	return w.ResponseWriter.Write(b) // 実際のレスポンスにも書き出す
}

func APILog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		// レスポンスライターを差し替え
		rcw := &responseCaptureWriter{
			ResponseWriter: res.Writer,
			body:           &bytes.Buffer{},
		}
		res.Writer = rcw

		// 処理実行
		err := next(c)

		// レスポンスボディから message を抽出
		var responseData map[string]interface{}
		var message string
		if err := json.Unmarshal(rcw.body.Bytes(), &responseData); err == nil {
			if msg, ok := responseData["message"].(string); ok {
				message = msg
			}
		}

		// 非同期保存
		go func() {
			log := models.APILog{
				Method:   req.Method,
				ClientIP: c.RealIP(),
				Path:     req.URL.Path,
				Message:  message,
			}
			db.DB.Create(&log)
		}()

		return err
	}
}


func BlockedIPLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		err := next(c)

		// チェック対象パスの配列を作成
		targetPaths := utils.RejectedPath

		// リクエストのパスが対象のパスの中にあるかをチェック
		for _, path := range targetPaths {
			if req.URL.Path == path {
				// 非同期でログ保存
				go func() {
					blockLog := models.BlockLog{
						ClientIP: c.RealIP(),
						Path:     req.URL.Path,
					}
					if err := db.DB.Create(&blockLog).Error; err != nil {
						log.Println("Error saving API log:", err)
					}
				}()
			}
		}

		return err
	}
}


