package webnotifications

import (
	"fmt"
	"net/http"
	"net/smtp"
	"net/url"
)

// NotifierType type of notifier
type NotifierType int

// Supported notifier types
const (
	Telegram NotifierType = iota
	Email
)

// NotifierConfig содержит конфигурацию для различных типов уведомлений
type NotifierConfig struct {
	// Telegram configs
	TGAPIKey    string
	TGChannelID string

	// Email configs
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
	ToEmail      string
}

// Notifier is an interface for sending notifications.
type Notifier interface {
	Notify(message string) error
}

// MultiNotifier поддерживает несколько типов уведомлений
type MultiNotifier struct {
	config    NotifierConfig
	notifiers []Notifier
}

// Добавьте методы валидации для каждого типа нотификатора
func isValidTelegramConfig(config NotifierConfig) bool {
	return config.TGAPIKey != "" && config.TGChannelID != ""
}

func isValidEmailConfig(config NotifierConfig) bool {
	return config.SMTPHost != "" &&
		config.SMTPPort > 0 &&
		config.SMTPUsername != "" &&
		config.SMTPPassword != "" &&
		config.FromEmail != "" &&
		config.ToEmail != ""
}

// NewMultiNotifier creates a new MultiNotifier with the given configuration and types.
// It will only add notifiers for the types that have valid configuration.
// Supported types are Telegram and Email.
func NewMultiNotifier(config NotifierConfig, types ...NotifierType) *MultiNotifier {
	mn := &MultiNotifier{
		config:    config,
		notifiers: make([]Notifier, 0),
	}

	for _, t := range types {
		switch t {
		case Telegram:
			if isValidTelegramConfig(config) {
				mn.notifiers = append(mn.notifiers, &telegramNotifier{
					apiKey:    config.TGAPIKey,
					channelID: config.TGChannelID,
				})
			}
		case Email:
			if isValidEmailConfig(config) {
				mn.notifiers = append(mn.notifiers, &emailNotifier{
					host:     config.SMTPHost,
					port:     config.SMTPPort,
					username: config.SMTPUsername,
					password: config.SMTPPassword,
					from:     config.FromEmail,
					to:       config.ToEmail,
				})
			}
		}
	}

	return mn
}

// Notify sends the given message to all notifiers.
func (mn *MultiNotifier) Notify(message string) error {
	var errors []error

	for _, n := range mn.notifiers {
		if err := n.Notify(message); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("notification errors: %v", errors)
	}
	return nil
}

// Telegram realisation
type telegramNotifier struct {
	apiKey    string
	channelID string
}

func (tn *telegramNotifier) Notify(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tn.apiKey)
	params := url.Values{}
	params.Add("chat_id", tn.channelID)
	params.Add("text", message)

	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return fmt.Errorf("telegram error: %v", err)
	}
	defer resp.Body.Close()
	return nil
}

// Email realisation
type emailNotifier struct {
	host     string
	port     int
	username string
	password string
	from     string
	to       string
}

func (en *emailNotifier) Notify(message string) error {
	addr := fmt.Sprintf("%s:%d", en.host, en.port)
	auth := smtp.PlainAuth("", en.username, en.password, en.host)

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Notification\r\n"+
		"\r\n"+
		"%s\r\n", en.to, message))

	err := smtp.SendMail(addr, auth, en.from, []string{en.to}, msg)
	if err != nil {
		return fmt.Errorf("email error: %v", err)
	}
	return nil
}
