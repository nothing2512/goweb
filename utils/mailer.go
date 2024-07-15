package utils

import (
	"bytes"
	"html/template"
	"os"

	"github.com/nothing2512/mailer"
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

	m, err := mailer.Init(
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_HOST"),
		os.Getenv("MAIL_PORT"),
	)
	if err != nil {
		return err
	}

	m.Recipients(target)
	m.Subject(title)
	m.SetHTML(body.String())

	return m.Send()
}
