package n

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func CreateProject(cfg *Config, slug []string) error {
	if err := cfg.EnsureDirs(); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	slugStr := strings.Join(slug, " ")
	if slugStr == "" {
		slugStr = "untitled"
	}

	timestamp := time.Now().Format("20060102-1504")
	filename := fmt.Sprintf("%s-%s.md", timestamp, strings.ReplaceAll(slugStr, " ", "-"))
	filepath := filepath.Join(cfg.ProjectsDir, filename)

	if err := os.WriteFile(filepath, []byte(ProjectNote(slugStr)), 0644); err != nil {
		return fmt.Errorf("failed to create project note: %w", err)
	}

	cmd := exec.Command(cfg.Editor, "+", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
