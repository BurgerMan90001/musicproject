package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"musicproject.com/config"
)

type Service struct {
	smtp smtp.Client
}

func New(cfg config.SMTP) (*Service, error) {
	smtpAuth := smtp.PlainAuth("", cfg.Email, cfg.Password, cfg.Host)
	client, err := smtp.Dial(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}
	// start tls
	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(&tls.Config{ServerName: cfg.Host}); err != nil {
			return nil, err
		}
	}
	if err := client.Auth(smtpAuth); err != nil {
		return nil, err
	}
	return &Service{}, nil
}
func (s *Service) SendOTP(otp string, recipient string) error {

	return nil
}
