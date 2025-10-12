package main

import (
	"fmt"
	"os"

	n "github.com/asurve/n"
	"github.com/spf13/cobra"
)

var cfg *n.Config

func main() {
	cfg = n.Load()

	rootCmd := &cobra.Command{
		Use:   "n",
		Short: "A simple note-taking CLI",
		Long:  "A simple and extensible note-taking CLI for daily notes, inbox captures, projects, and search.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return n.OpenDaily(cfg)
		},
	}

	dailyCmd := &cobra.Command{
		Use:   "daily",
		Short: "Open today's daily note",
		RunE: func(cmd *cobra.Command, args []string) error {
			return n.OpenDaily(cfg)
		},
	}

	inboxCmd := &cobra.Command{
		Use:   "inbox [text...]",
		Short: "Quick append to inbox with timestamp",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return n.AppendInbox(cfg, args)
		},
	}

	newCmd := &cobra.Command{
		Use:   "new [slug...]",
		Short: "Create a new project note",
		RunE: func(cmd *cobra.Command, args []string) error {
			return n.CreateProject(cfg, args)
		},
	}

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search notes with ripgrep and fzf",
		RunE: func(cmd *cobra.Command, args []string) error {
			return n.Search(cfg)
		},
	}

	rootCmd.AddCommand(dailyCmd, inboxCmd, newCmd, searchCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
