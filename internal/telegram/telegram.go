package telegram

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/vova4o/webnotifications/models"
	"golang.org/x/time/rate"
)

// Global message limitter: 25 requests per second
var (
	// Global rate limiter: 25 запросов в секунду
	globalLimiter = rate.NewLimiter(rate.Limit(25), 25)

	// Map of channel limiters
	channelLimiters   = make(map[string]*rate.Limiter)
	channelLimitersMu sync.Mutex
)

// getChannelLimiter returns a rate limiter for the given channel ID, or creates a new one
func getChannelLimiter(channelID string) *rate.Limiter {
	channelLimitersMu.Lock()
	defer channelLimitersMu.Unlock()

	limiter, exists := channelLimiters[channelID]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(1), 1) // 1 запрос в секунду, бурст 1
		channelLimiters[channelID] = limiter
	}

	return limiter
}

// telegramNotifier structure implements Notifier interface for Telegram
type telegramNotifier struct {
	apiKey    string
	channelID string
}

// Notify sends a Telegram notification with context support
func (tn *telegramNotifier) Notify(ctx context.Context, message string) error {
	if err := globalLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("global rate limiter error: %v", err)
	}

	channelLimiter := getChannelLimiter(tn.channelID)

	if err := channelLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("channel rate limiter error: %v", err)
	}

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

// Create implements a factory method to create a notifier
func Create(config models.NotifierConfig) models.Notifier {
	return &telegramNotifier{
		apiKey:    config.TGAPIKey,
		channelID: config.TGChannelID,
	}
}

// IsValidTelegramConfig checks the validity of the Telegram configuration
func IsValidTelegramConfig(config models.NotifierConfig) bool {
	return config.TGAPIKey != "" && config.TGChannelID != ""
}
