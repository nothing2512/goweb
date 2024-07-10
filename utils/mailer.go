package utils

import (
	"bytes"
	"html/template"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(title, target, view string, data interface{}) error {
	t, err := template.ParseFiles(view)
	if err != nil {
		return err
	}
	var body bytes.Buffer
	if data != nil {
		_ = t.Execute(&body, data)
	} else {
		_ = t.Execute(&body, struct{}{})
	}

	mail := gomail.NewMessage()
	mail.FormatAddress(target, "")

	mail.SetHeader("From", os.Getenv("MAIL_USER"))
	mail.SetHeader("To", target)
	mail.SetHeader("Subject", title)
	mail.SetHeader("List-Unsubscribe", os.Getenv("MAIL_USER"))
	mail.SetBody("text/html", body.String())

	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	dialer := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		port,
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASS"),
	)

	err = dialer.DialAndSend(mail)
	if err != nil {
		return err
	}
	return nil
}
