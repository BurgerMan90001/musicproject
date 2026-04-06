package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"musicproject.com/internal/config/secrets"
)

type Service struct {
	smtp smtp.Client
}

func New(ctx context.Context, sm secrets.Manager) (*Service, error) {
	s, err := secrets.GetSecrets(ctx, sm, "SMTP_EMAIL",
		"SMTP_PASSWORD", "SMTP_HOST", "SMTP_PORT")
	if err != nil {
		return nil, err
	}

	smtpAuth := smtp.PlainAuth("", s[0], s[1], s[2])
	client, err := smtp.Dial(fmt.Sprintf("%s:%s", s[2], s[3]))
	if err != nil {
		return nil, fmt.Errorf("Email service dial: %v", err)
	}
	// start tls
	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(&tls.Config{ServerName: s[2]}); err != nil {
			return nil, fmt.Errorf("Email service startTLS: %w", err)
		}
	}
	// authenticate
	if err := client.Auth(smtpAuth); err != nil {
		return nil, fmt.Errorf("Email service auth: %v", err)
	}
	return &Service{}, nil
}

func (s *Service) SendWelcome() {
	
}
func (s *Service) SendOTP(otp string, recipient string) error {
	return nil
}
