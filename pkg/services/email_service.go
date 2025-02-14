package services

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

func NewEmailService(smtpHost, smtpPort, username, password string) *EmailService {
	return &EmailService{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", e.username, e.password, e.smtpHost)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.username, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
