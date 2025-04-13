package cmd

import (
	"github.com/giovanni-gava/argocd-pipeline-trigger/internal/argocd"

	"fmt"

	"github.com/spf13/cobra"
)

var (
	appName  string
	server   string
	user     string
	pass     string
	insecure bool
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Triggers ArgoCD app sync",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := argocd.Config{
			App:      appName,
			Server:   server,
			Username: user,
			Password: pass,
			Insecure: insecure,
		}
		if err := argocd.SyncApp(cfg); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		} else {
			fmt.Println("✅ Sync triggered successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVar(&appName, "app", "", "ArgoCD app name (required)")
	syncCmd.Flags().StringVar(&server, "server", "", "ArgoCD server URL (required)")
	syncCmd.Flags().StringVar(&user, "username", "", "Username")
	syncCmd.Flags().StringVar(&pass, "password", "", "Password")
	syncCmd.Flags().BoolVar(&insecure, "insecure", false, "Skip TLS verify")

	syncCmd.MarkFlagRequired("app")
	syncCmd.MarkFlagRequired("server")
}
