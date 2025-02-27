package email

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/vova4o/webnotifications/models"
)

// emailNotifier structure implements Notifier interface for Email
type emailNotifier struct {
	host     string
	port     int
	username string
	password string
	from     string
	to       string
}

// Notify sends an Email notification with context support
func (en *emailNotifier) Notify(ctx context.Context, message string) error {
	doneCh := make(chan error, 1)

	go func() {
		addr := fmt.Sprintf("%s:%d", en.host, en.port)
		auth := smtp.PlainAuth("", en.username, en.password, en.host)

		msg := []byte(fmt.Sprintf("To: %s\r\n"+
			"Subject: Notification\r\n"+
			"\r\n"+
			"%s\r\n", en.to, message))

		err := smtp.SendMail(addr, auth, en.from, []string{en.to}, msg)
		doneCh <- err
	}()

	select {
	case err := <-doneCh:
		if err != nil {
			return fmt.Errorf("email error: %v", err)
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("email sending canceled: %v", ctx.Err())
	}
}

// Create creates a new Email notifier
func Create(config models.NotifierConfig) models.Notifier {
	return &emailNotifier{
		host:     config.SMTPHost,
		port:     config.SMTPPort,
		username: config.SMTPUsername,
		password: config.SMTPPassword,
		from:     config.FromEmail,
		to:       config.ToEmail,
	}
}

// IsValidEmailConfig checks if the Email notifier config is valid
func IsValidEmailConfig(config models.NotifierConfig) bool {
	return config.SMTPHost != "" &&
		config.SMTPPort > 0 &&
		config.SMTPUsername != "" &&
		config.SMTPPassword != "" &&
		config.FromEmail != "" &&
		config.ToEmail != ""
}
