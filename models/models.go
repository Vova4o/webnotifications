package models

import "context"

// NotifierType notates the type of notifier
type NotifierType int

// supported types of notifications
const (
	Telegram NotifierType = iota
	Email
)

// NotifierConfig is a configuration for a notifier.
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
	Notify(ctx context.Context, message string) error
}
