// internal/notifier/telegram.go
package notifier

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type TelegramNotifier struct {
	token  string
	chatID string
}

func NewTelegramNotifier(token, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		token:  token,
		chatID: chatID,
	}
}

func (t *TelegramNotifier) Notify(ctx context.Context, title, message string, success bool) error {
	emoji := "✅"
	if !success {
		emoji = "❌"
	}

	formatted := fmt.Sprintf("%s *%s*\n%s", emoji, title, message)
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.token)

	values := url.Values{}
	values.Set("chat_id", t.chatID)
	values.Set("text", formatted)
	values.Set("parse_mode", "Markdown")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("telegram returned status %d", resp.StatusCode)
	}
	return nil
}
