package notes

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/asurve/n/internal/config"
)

func AppendInbox(cfg *config.Config, text []string) error {
	if err := cfg.EnsureDirs(); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	filepath := filepath.Join(cfg.InboxDir, "inbox.md")

	timestamp := time.Now().Format("2006-01-02 15:04")
	entry := fmt.Sprintf("- [%s] %s\n", timestamp, strings.Join(text, " "))

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open inbox: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(entry); err != nil {
		return fmt.Errorf("failed to write to inbox: %w", err)
	}

	cmd := exec.Command(cfg.Editor, "+", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
