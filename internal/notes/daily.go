package notes

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/asurve/n/internal/config"
	"github.com/asurve/n/internal/templates"
)

func OpenDaily(cfg *config.Config) error {
	if err := cfg.EnsureDirs(); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	filename := time.Now().Format("2006-01-02") + ".md"
	filepath := filepath.Join(cfg.DailyDir, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if err := os.WriteFile(filepath, []byte(templates.DailyNote()), 0644); err != nil {
			return fmt.Errorf("failed to create daily note: %w", err)
		}
	}

	cmd := exec.Command(cfg.Editor, "+", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
