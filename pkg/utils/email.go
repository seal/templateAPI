package utils

import (
	"bytes"
	"crypto/tls"
	"errors"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/k3a/html2text"
	"github.com/seal/templateapi/pkg/models"

	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

// Verification email
func SendEmail(user *models.User, data *EmailData) {

	// Sender data.
	from := EnvVariable("EmailFrom")
	smtpPass := EnvVariable("SMTPPass")
	smtpUser := EnvVariable("SMTPUser")
	to := user.Email
	smtpHost := EnvVariable("SMTPHost")
	smtpPort, err := strconv.Atoi(EnvVariable("SMTPPort"))
	if err != nil {
		Error(errors.New("Error converting port to int, env variable" + err.Error()))
	}

	var body bytes.Buffer

	template, err := ParseTemplateDir("pkg/templates/")
	if err != nil {
		Error(errors.New("Could not parse template" + err.Error()))
		return
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}
