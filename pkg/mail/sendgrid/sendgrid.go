package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	registerDto "online-course/internal/register/dto"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail interface {
	SendVerificationEmail(toEmail string, dto registerDto.CreateEmailVerification)
}

type MailImpl struct {
}

func (mi *MailImpl) sendMail(toEmail string, result string, subject string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER_NAME"), os.Getenv("MAIL_SENDER_NAME"))
	to := mail.NewEmail(toEmail, toEmail)

	messages := mail.NewSingleEmail(from, subject, to, "", result)

	client := sendgrid.NewSendClient(os.Getenv("MAIL_KEY"))

	resp, err := client.Send(messages)

	if err != nil {
		fmt.Println(err)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp)
	} else {
		fmt.Println("email berhasil dikirim ke %s", toEmail)
	}
}

// SendVerificationEmail implements Mail
func (mi *MailImpl) SendVerificationEmail(toEmail string, dto registerDto.CreateEmailVerification) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/emails/verification_email.html")

	result, err := ParseTemplate(templateFile, dto)

	if err != nil {
		fmt.Print(err)
	}

	mi.sendMail(toEmail, result, dto.SUBJECT)
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewMail() Mail {
	return &MailImpl{}
}
