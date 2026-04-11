package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"musicproject.com/internal/config/secrets"
)

type Sender interface {
	Send(ctx context.Context, to, subject, body string) error
}

type Service struct {
	smtp smtp.Client
}

func New(ctx context.Context, sm secrets.Manager) (*Service, error) {
	secretList, err := secrets.GetSecrets(ctx, sm, "SMTP_EMAIL",
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
	return &Service{}, nil
}

// TODO MAYBE
func (s *Service) SendWelcome() {

}
func (s *Service) SendOTP(otp string, recipient string) error {
	return nil
}
