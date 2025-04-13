// internal/notifier/notifier.go
package notifier

import (
	"context"
	"fmt"
	"os"
	"time"
)

type Notifier interface {
	Notify(ctx context.Context, title, message string, success bool) error
}

// Config holds configuration for available notifiers
type Config struct {
	EnableSlack    bool
	SlackWebhook   string
	EnableTelegram bool
	TelegramToken  string
	TelegramChatID string
}

// NewFromEnv initializes the notifier from environment variables
func NewFromEnv() Notifier {
	cfg := Config{
		EnableSlack:    os.Getenv("ENABLE_SLACK") == "true",
		SlackWebhook:   os.Getenv("SLACK_WEBHOOK_URL"),
		EnableTelegram: os.Getenv("ENABLE_TELEGRAM") == "true",
		TelegramToken:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChatID: os.Getenv("TELEGRAM_CHAT_ID"),
	}

	multi := MultiNotifier{}
	if cfg.EnableSlack && cfg.SlackWebhook != "" {
		multi = append(multi, NewSlackNotifier(cfg.SlackWebhook))
	}
	if cfg.EnableTelegram && cfg.TelegramToken != "" && cfg.TelegramChatID != "" {
		multi = append(multi, NewTelegramNotifier(cfg.TelegramToken, cfg.TelegramChatID))
	}

	return multi
}

// MultiNotifier allows multiple notifiers to run in parallel
type MultiNotifier []Notifier

func (m MultiNotifier) Notify(ctx context.Context, title, message string, success bool) error {
	errs := []error{}
	for _, n := range m {
		ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		if err := n.Notify(ctxTimeout, title, message, success); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("notification errors: %v", errs)
	}
	return nil
}
