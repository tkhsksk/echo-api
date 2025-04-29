package middlewares

import (
	"bytes"
	"html/template"
	"os"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(address string, name string) error {
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
		log.Println("Error parsing template:", err)
		return err
	}

	var body bytes.Buffer
	data := struct {
		Name string
	}{
		Name: name,
	}

	if err := t.Execute(&body, data); err != nil {
		log.Println("Error executing template:", err)
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

	log.Println("HTML email sent successfully!")
	return nil
}
