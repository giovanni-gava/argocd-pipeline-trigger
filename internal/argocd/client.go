package argocd

import (
	"fmt"
	"os/exec"
)

func SyncApp(cfg Config) error {
	loginArgs := []string{
		"login", cfg.Server,
		"--username", cfg.Username,
		"--password", cfg.Password,
	}
	if cfg.Insecure {
		loginArgs = append(loginArgs, "--insecure")
	}

	syncArgs := []string{
		"app", "sync", cfg.App,
	}

	// Login
	if out, err := exec.Command("argocd", loginArgs...).CombinedOutput(); err != nil {
		return fmt.Errorf("login failed: %s", string(out))
	}

	// Sync
	if out, err := exec.Command("argocd", syncArgs...).CombinedOutput(); err != nil {
		return fmt.Errorf("sync failed: %s", string(out))
	}

	return nil
}
