package services

import (
	"fmt"
	"net/smtp"
	// "net/smtp"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

type EmailSender interface {
	Send(to string, subject string, plainText string, htmlBody string) error
}

type SendGridConfig struct {
	EmailClient *sendgrid.Client
	EmailFrom   string
}

func NewSendGridConfig(apiKey string, emailTo string) *SendGridConfig {
	client := sendgrid.NewSendClient(apiKey)
	return &SendGridConfig{
		EmailClient: client,
		EmailFrom:   emailTo,
	}
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func NewSMTPConfig(host string, port string, username string, password string) *SMTPConfig {
	smtpConfig := &SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
	return smtpConfig
}

var sendOTPService *SendOTPService

type SendOTPService struct {
	sendGridConfig *SendGridConfig
	smtpConfig     *SMTPConfig
	log            *logrus.Logger
}

func NewSendOTPService(sendGridConfig *SendGridConfig, smtpConfig *SMTPConfig, log *logrus.Logger) {
	sendOTPService = &SendOTPService{
		sendGridConfig: sendGridConfig,
		smtpConfig:     smtpConfig,
		log:            log,
	}
}

func (s *SMTPConfig) Send(to string, subject string, plainText string, htmlBody string) error {
	from := s.Username
	displayName := "Healthmate App"
	msg := "MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		fmt.Sprintf("From: \"%s\" <%s>\r\n", displayName, from) +
		fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("Subject: %s\r\n\r\n", subject) +
		htmlBody

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)

	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("SMTP error: %w", err)
	}
	return nil
}

func (s *SendGridConfig) Send(to string, subject string, plainText string, htmlBody string) error {
	from := mail.NewEmail("Healthmate App", s.EmailFrom)
	toEmail := mail.NewEmail("User", to)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.Subject = subject
	message.AddContent(mail.NewContent("text/plain", plainText))
	message.AddContent(mail.NewContent("text/html", htmlBody))

	p := mail.NewPersonalization()
	p.AddTos(toEmail)
	message.AddPersonalizations(p)

	response, err := s.EmailClient.Send(message)
	if err != nil {
		return fmt.Errorf("SendGrid error: %w", err)
	}
	if response.StatusCode >= 400 {
		return fmt.Errorf("SendGrid failed with status: %d - %s", response.StatusCode, response.Body)
	}
	return nil

}

func SendOTPEmail(to string, otpCode string) error {
	subject := "Your OTP Code"
	plainText := fmt.Sprintf(
		"Your OTP code is: %s. Please do not share this code.\nThis code will expire in 10 minutes.",
		otpCode,
	)

	htmlBody := fmt.Sprintf(
		`<p>Your OTP code is: <strong>%s</strong>. Please do not share this code.</p>
     <p><em>This code will expire in 10 minutes.</em></p>`,
		otpCode,
	)

	// Gửi bằng SendGrid trước
	err := sendOTPService.sendGridConfig.Send(to, subject, plainText, htmlBody)
	if err != nil {
		sendOTPService.log.Warnf("SendGrid failed: %v. Falling back to SMTP...", err)

		// Nếu SendGrid lỗi, thử gửi bằng SMTP
		errSMTP := sendOTPService.smtpConfig.Send(to, subject, plainText, htmlBody)
		if errSMTP != nil {
			sendOTPService.log.Errorf("SMTP also failed: %v", errSMTP)
			return errSMTP
		}
		sendOTPService.log.Infof("Email sent successfully via SMTP to %s", to)
		return nil
	}

	sendOTPService.log.Infof("Email sent successfully via SendGrid to %s", to)
	return nil
}
