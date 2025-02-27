package main

import (
    "context"
    "fmt"
    "time"

    "github.com/vova4o/webnotifications"
    "github.com/vova4o/webnotifications/models"
)

func main() {
    // Настройка конфигурации нотификаторов
    config := models.NotifierConfig{
        TGAPIKey:    "your-telegram-bot-api-key",
        TGChannelID: "your-telegram-channel-id",
        SMTPHost:     "smtp.example.com",
        SMTPPort:     587,
        SMTPUsername: "your-email@example.com",
        SMTPPassword: "your-email-password",
        FromEmail:    "your-email@example.com",
        ToEmail:      "recipient@example.com",
    }

    // Создание MultiNotifier с поддержкой Telegram и Email
    notifier := webnotifications.NewMultiNotifier(config, models.Telegram, models.Email)

    // Простой вызов
    err := notifier.Notify("Hello, this is a test notification!")
    if err != nil {
        fmt.Printf("Error sending notification: %v\n", err)
    }

    // Использование с контекстом и таймаутом
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = notifier.NotifyWithContext(ctx, "Hello with timeout!")
    if err != nil {
        fmt.Printf("Error sending notification with context: %v\n", err)
    }
}