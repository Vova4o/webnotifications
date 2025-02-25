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
go get github.com/yourusername/webnotifications
```

## Usage
### Importing the package
```go
import "github.com/yourusername/webnotifications"
```

### Setting up NotifierConfig
```go
config := webnotifications.NotifierConfig{
    TGAPIKey:    "your-telegram-bot-api-key",
    TGChannelID: "your-telegram-channel-id",
    SMTPHost:     "smtp.example.com",
    SMTPPort:     587,
    SMTPUsername: "your-email@example.com",
    SMTPPassword: "your-email-password",
    FromEmail:    "your-email@example.com",
    ToEmail:      "recipient@example.com",
}
```

### Creating and Using MultiNotifier
```go
notifier := webnotifications.NewMultiNotifier(config, webnotifications.Telegram, webnotifications.Email)

err := notifier.Notify("Hello, this is a test notification!")
if err != nil {
    fmt.Println("Error sending notification:", err)
}
```

## Environment Variables (Optional)
Instead of hardcoding credentials, consider using environment variables:
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
Then retrieve them in your Go code:
```go
import "os"

config := webnotifications.NotifierConfig{
    TGAPIKey:    os.Getenv("TG_API_KEY"),
    TGChannelID: os.Getenv("TG_CHANNEL_ID"),
    SMTPHost:     os.Getenv("SMTP_HOST"),
    SMTPPort:     587,
    SMTPUsername: os.Getenv("SMTP_USERNAME"),
    SMTPPassword: os.Getenv("SMTP_PASSWORD"),
    FromEmail:    os.Getenv("FROM_EMAIL"),
    ToEmail:      os.Getenv("TO_EMAIL"),
}
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

