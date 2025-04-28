package handlers

import (
	"gopkg.in/gomail.v2"

	"os"
	"log"
	"strconv"
	"github.com/joho/godotenv"
)

func sendMail(address string) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	from := os.Getenv("MAIL_FROM")
	smtp := os.Getenv("MAIL_SMTP")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	user := os.Getenv("MAIL_USER")
	pass := os.Getenv("MAIL_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", "\"api admin\" <" + from + ">")
	m.SetHeader("To", address)
	m.SetHeader("Subject", "subject")
	m.SetBody("text/plain", "This is the email body.")

	d := gomail.NewDialer(smtp, port, user, pass)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email: ", err)
		return err
	}
	log.Println("Email sent successfully!")
	return nil
}
