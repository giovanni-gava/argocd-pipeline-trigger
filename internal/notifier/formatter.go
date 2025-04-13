package notifier

import (
	"fmt"
	"time"
)

// FormatMessage gera uma mensagem amigável e visual para Slack/Telegram
func FormatMessage(app string, duration time.Duration, err error) (title string, message string, success bool) {
	success = err == nil
	title = "Sync Triggered"

	if !success {
		title = "Sync Failed"
	}

	message = fmt.Sprintf(
		"👨🏻‍💻 *Giovanni Colognesi - ArgoCD Trigger*\n"+
			"🔹 App: `%s`\n"+
			"⏱️ Duration: `%s`",
		app,
		duration.String(),
	)

	if err != nil {
		message += fmt.Sprintf("\n❗ Error: `%s`", err.Error())
	}

	return
}
