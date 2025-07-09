package models

import "github.com/go-mail/mail/v2"

const (
	DefaultSender = "support@lenslocked.com"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type EmailService struct {
	DefaultSender string

	dialer *mail.Dialer
}

func NewEmailService(cfg SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password),
	}
	return &es
}
