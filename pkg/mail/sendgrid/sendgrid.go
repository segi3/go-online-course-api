package mail

import (
	"bytes"
	"fmt"
	"html/template"
	emailRegisterDto "online-course/internal/register/dto"
	"os"
	"path/filepath"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail interface {
	SendVerificationEmail(toEmail string, dto emailRegisterDto.CreateEmailVerification)
}

type MailImpl struct {
}

func (mi MailImpl) sendMail(toEmail string, result string, subject string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER"), os.Getenv("MAIL_SENDER"))
	to := mail.NewEmail(toEmail, toEmail)

	message := mail.NewSingleEmail(from, subject, to, "", result)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)

	if err != nil {
		fmt.Println(err)
	} else if response.StatusCode != 200 {
		fmt.Println(response)
	} else {
		fmt.Println("Email successfullt send")
	}
}

// SendVerificationEmail implements Mail
func (mi *MailImpl) SendVerificationEmail(toEmail string, dto emailRegisterDto.CreateEmailVerification) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/emails/verification_email.html")

	result, err := Parsetemplate(templateFile, dto)

	if err != nil {
		fmt.Println(err)
	}

	mi.sendMail(toEmail, result, dto.SUBJECT)
}

// parse html template
func Parsetemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "nil", err
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, data); err != nil {
		return "nil", err
	}

	return buf.String(), nil
}

func NewMail() Mail {
	return &MailImpl{}
}
