package mailer

import (
	"bytes"
	"html/template"
	"os"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendPasscodeMail(address string, code string, user_id uint, passcode_id uint) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	from := os.Getenv("MAIL_FROM")
	smtp := os.Getenv("MAIL_SMTP")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	user := os.Getenv("MAIL_USER")
	pass := os.Getenv("MAIL_PASS")

	// テンプレート読み込みとデータの埋め込み
	t, err := template.ParseFiles("templates/email.html")
	if err != nil {
		log.Println("テンプレートエラー:", err)
		return err
	}

	var body bytes.Buffer
	data := struct {
		UserId uint
		CodeId uint
		Code   string
		Url    string
	}{
		UserId: user_id,
		CodeId: passcode_id,
		Code:   code,
		Url:    os.Getenv("API_URL"),
	}

	if err := t.Execute(&body, data); err != nil {
		log.Println("テンプレート内部エラー:", err)
		return err
	}

	// メール作成
	m := gomail.NewMessage()
	m.SetHeader("From", "\"API service\" <"+from+">")
	m.SetHeader("To", address)
	m.SetHeader("Subject", "HTMLメールの件名")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(smtp, port, user, pass)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Println("emailの送信に成功しました")
	return nil
}
