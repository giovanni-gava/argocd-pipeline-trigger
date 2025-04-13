package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/giovanni-gava/argocd-pipeline-trigger/internal/notifier"
)

func main() {
	ctx := context.Background()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if token == "" || chatID == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID must be set")
	}

	tg := notifier.NewTelegramNotifier(token, chatID)

	title, msg, success := notifier.FormatMessage("test-app", 1300*time.Millisecond, nil)
	if err := tg.Notify(ctx, title, msg, success); err != nil {
		log.Fatalf("❌ Error sending notification: %v", err)
	}
	log.Println("✅ Notification sent successfully")
}
