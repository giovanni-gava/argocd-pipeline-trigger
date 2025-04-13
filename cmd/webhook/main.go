package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
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
	http.HandleFunc("/sync", loggingMiddleware(syncHandler))
	log.Printf("🌐 Webhook receiver listening on %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func syncHandler(w http.ResponseWriter, r *http.Request) {
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

	err := triggerArgoCDSync(ctx, payload.App)
	if err != nil {
		log.Printf("❌ Failed to sync app %q: %v", payload.App, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Successfully triggered sync for app: %s", payload.App)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
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
