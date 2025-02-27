# WebNotifications

WebNotifications is a Go package that provides a flexible notification system supporting multiple notification channels, including Telegram and Email. It allows sending messages through various channels by configuring valid notifier settings.

## Features

- Supports Telegram notifications
- Supports Email notifications via SMTP
- Allows configuring multiple notifiers at once
- Validates configuration before adding a notifier

## Installation

To install WebNotifications, use:

```sh
go get github.com/vova4o/webnotifications
```

## Importing

To import WebNotifications in your Go project, use:

```go
import (
    "github.com/vova4o/webnotifications"
    "github.com/vova4o/webnotifications/models"
)
```

## Configuration

To configure WebNotifications, create a `NotifierConfig` struct with the necessary settings:

```go
config := models.NotifierConfig{
    TGAPIKey:    "your-telegram-bot-api-key",
    TGChannelID: "your-telegram-channel-id",
    SMTPHost:     "smtp.example.com",
    SMTPPort:     587,
    SMTPUsername: "your-email@example.com",
    SMTPPassword: "your-email-password", // make sure to check your email provider's security settings
    // some email providers require you to use app passwords instead of your account password
    FromEmail:    "your-email@example.com",
    ToEmail:      "recipient@example.com",
}
```

## Usage

To use WebNotifications, create a new notifier and send notifications:

```go
notifier := webnotifications.NewMultiNotifier(config, models.Telegram, models.Email)

// Simple call
err := notifier.Notify("Hello, this is a test notification!")

// Using with context and timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
err := notifier.NotifyWithContext(ctx, "Hello with timeout!")
```

## Examples

To run the examples, navigate to the `examples` directory and run the `main.go` file:

```sh
cd examples
go run main.go
```

## Environment Variables (Optional)

Instead of hardcoding your configuration, you can use environment variables to set your notifier settings:

```sh
export TG_API_KEY="your-telegram-bot-api-key"
export TG_CHANNEL_ID="your-telegram-channel-id"
export SMTP_HOST="smtp.example.com"
export SMTP_PORT="587"
export SMTP_USERNAME="your-email@example.com"
export SMTP_PASSWORD="your-email-password"
export FROM_EMAIL="your-email@example.com"
export TO_EMAIL="recipient@example.com"
```

To retrieve the environment variables in your Go code, use:

```go
config := models.NotifierConfig{
    TGAPIKey:    os.Getenv("TG_API_KEY"),
    TGChannelID: os.Getenv("TG_CHANNEL_ID"),
    SMTPHost:    os.Getenv("SMTP_HOST"),
    SMTPPort:    os.Getenv("SMTP_PORT"),
    SMTPUsername: os.Getenv("SMTP_USERNAME"),
    SMTPPassword: os.Getenv("SMTP_PASSWORD"),
    FromEmail:   os.Getenv("FROM_EMAIL"),
    ToEmail:     os.Getenv("TO_EMAIL"),
}
```
