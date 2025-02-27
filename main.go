package webnotifications

import (
	"context"
	"fmt"

	"github.com/vova4o/webnotifications/internal/email"
	"github.com/vova4o/webnotifications/internal/telegram"
	"github.com/vova4o/webnotifications/models"
)

// MultiNotifier is a notifier that sends messages to multiple notifiers.
type MultiNotifier struct {
	config    models.NotifierConfig
	notifiers []models.Notifier
}

// NewMultiNotifier creates a new MultiNotifier with the given configuration and types.
// It will only add notifiers for the types that have valid configuration.
// Supported types are Telegram and Email.
func NewMultiNotifier(config models.NotifierConfig, types ...models.NotifierType) *MultiNotifier {
	mn := &MultiNotifier{
		config:    config,
		notifiers: make([]models.Notifier, 0),
	}

	for _, t := range types {
		switch t {
		case models.Telegram:
			if telegram.IsValidTelegramConfig(config) {
				// use the telegram package to create a notifier
				mn.notifiers = append(mn.notifiers, telegram.Create(config))
			}
		case models.Email:
			if email.IsValidEmailConfig(config) {
				// use the email package to create a notifier
				mn.notifiers = append(mn.notifiers, email.Create(config))
			}
		}
	}

	return mn
}

// Notify sends the given message to all notifiers.
func (mn *MultiNotifier) Notify(message string) error {
	// Create a background context
	ctx := context.Background()
	return mn.NotifyWithContext(ctx, message)
}

// NotifyWithContext sends the given message to all notifiers with a specific context.
func (mn *MultiNotifier) NotifyWithContext(ctx context.Context, message string) error {
	var errors []error

	for _, n := range mn.notifiers {
		if err := n.Notify(ctx, message); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("notification errors: %v", errors)
	}
	return nil
}
