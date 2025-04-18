// cmd/webhook/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/giovanni-gava/argocd-pipeline-trigger/internal/notifier"
)

const (
	defaultAddr = ":8080"
	timeout     = 10 * time.Second
)

type Payload struct {
	App string `json:"app"`
}

func main() {
	addr := getEnv("ADDR", defaultAddr)
	n := notifier.NewFromEnv()
	notifier.RegisterMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/sync", loggingMiddleware(withAuth(syncHandler(n))))

	log.Printf("🌐 Webhook receiver listening on %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expectedToken := os.Getenv("AUTH_TOKEN")
		if expectedToken == "" {
			next(w, r)
			return
		}

		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") || strings.TrimPrefix(auth, "Bearer ") != expectedToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func syncHandler(n notifier.Notifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		var payload Payload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if payload.App == "" {
			http.Error(w, "App name required", http.StatusBadRequest)
			return
		}

		start := time.Now()
		err := triggerArgoCDSync(ctx, payload.App)
		duration := time.Since(start)

		title, msg, success := notifier.FormatMessage(payload.App, duration, err)
		_ = n.Notify(ctx, title, msg, success)

		status := "success"
		if err != nil {
			status = "failure"
			notifier.SyncTotal.WithLabelValues(payload.App, status).Inc()
			notifier.SyncDuration.WithLabelValues(payload.App).Observe(duration.Seconds())
			log.Printf("❌ Failed to sync app %q: %v", payload.App, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		notifier.SyncTotal.WithLabelValues(payload.App, status).Inc()
		notifier.SyncDuration.WithLabelValues(payload.App).Observe(duration.Seconds())

		log.Printf("✅ Successfully triggered sync for app: %s", payload.App)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

func triggerArgoCDSync(ctx context.Context, app string) error {
	cmd := exec.CommandContext(ctx, "argocd", "app", "sync", app)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %v\noutput: %s", err, string(out))
	}
	return nil
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.Printf("➡️ %s %s [%s]", r.Method, r.URL.Path, duration)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
