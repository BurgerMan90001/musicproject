package email

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"

	"musicproject.com/internal/config/secrets"
)

type Email struct {
	subject string
	to      []string
	from    string
	headers string
	body    []byte
}

type Sender interface {
	Send(ctx context.Context, to, subject, body string) error
}

type Service struct {
	client *smtp.Client
}

func New() (*Service, error) {
	secretList, err := secrets.GetEnv("SMTP_EMAIL",
		"SMTP_PASSWORD", "SMTP_HOST", "SMTP_PORT")
	if err != nil {
		return nil, err
	}

	smtpAuth := smtp.PlainAuth("", secretList["SMTP_EMAIL"],
		secretList["SMTP_PASSWORD"], secretList["SMTP_HOST"])

	addr := fmt.Sprintf("%s:%s", secretList["SMTP_HOST"], secretList["SMTP_PORT"])

	client, err := smtp.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("Email service dial: %v", err)
	}
	// start tls connection
	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(&tls.Config{ServerName: secretList["SMTP_HOST"]}); err != nil {
			return nil, fmt.Errorf("Email service startTLS: %w", err)
		}
	}
	// authenticate
	if err := client.Auth(smtpAuth); err != nil {
		return nil, fmt.Errorf("Email service auth: %v", err)
	}
	//client.Quit()

	return &Service{client: client}, nil
}

func (s *Service) SendMail(email *Email) error {
	if email == nil {
		return errors.New("email is nil")
	}
	// Set the sender
	if err := s.client.Mail(email.from); err != nil {
		return err
	}
	// Set the recipents
	for _, subject := range email.to {
		if err := s.client.Rcpt(subject); err != nil {
			return err
		}
	}
	w, err := s.client.Data()
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "Subject: %v\n%v\n\n%v",
		email.subject, email.headers, string(email.body)); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
func (s *Service) sendMailTemplate(email *Email, templatePath string) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("template parse: %w", err)
	}
	t.Execute(&body, struct{ Name string }{Name: "Goob"})
	email.body = body.Bytes()
	email.headers = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	return s.SendMail(email)
}
func (s *Service) Cleanup() {
	s.client.Quit()
}

// TODO MAYBE
func (s *Service) SendWelcome(from string) {

}
func (s *Service) SendOTP(otp string, from string) error {
	return nil
}
