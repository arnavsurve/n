package main

import (
	"fmt"
	"os"

	"github.com/asurve/n/internal/config"
	"github.com/asurve/n/internal/notes"
	"github.com/spf13/cobra"
)

var cfg *config.Config

func main() {
	cfg = config.Load()

	rootCmd := &cobra.Command{
		Use:   "n",
		Short: "A simple note-taking CLI",
		Long:  "A simple and extensible note-taking CLI for daily notes, inbox captures, projects, and search.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return notes.OpenDaily(cfg)
		},
	}

	dailyCmd := &cobra.Command{
		Use:   "daily",
		Short: "Open today's daily note",
		RunE: func(cmd *cobra.Command, args []string) error {
			return notes.OpenDaily(cfg)
		},
	}

	inboxCmd := &cobra.Command{
		Use:   "inbox [text...]",
		Short: "Quick append to inbox with timestamp",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return notes.AppendInbox(cfg, args)
		},
	}

	newCmd := &cobra.Command{
		Use:   "new [slug...]",
		Short: "Create a new project note",
		RunE: func(cmd *cobra.Command, args []string) error {
			return notes.CreateProject(cfg, args)
		},
	}

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search notes with ripgrep and fzf",
		RunE: func(cmd *cobra.Command, args []string) error {
			return notes.Search(cfg)
		},
	}

	rootCmd.AddCommand(dailyCmd, inboxCmd, newCmd, searchCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
