package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	NotesDir    string
	Editor      string
	DailyDir    string
	InboxDir    string
	ProjectsDir string
	ScratchDir  string
}

func Load() *Config {
	homeDir, _ := os.UserHomeDir()
	notesDir := filepath.Join(homeDir, "notes")

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nvim"
	}

	return &Config{
		NotesDir:    notesDir,
		Editor:      editor,
		DailyDir:    filepath.Join(notesDir, "daily"),
		InboxDir:    filepath.Join(notesDir, "inbox"),
		ProjectsDir: filepath.Join(notesDir, "projects"),
		ScratchDir:  filepath.Join(notesDir, "scratch"),
	}
}

func (c *Config) EnsureDirs() error {
	dirs := []string{
		c.DailyDir,
		c.InboxDir,
		c.ProjectsDir,
		c.ScratchDir,
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}
