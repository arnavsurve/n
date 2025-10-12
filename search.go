package n

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Search(cfg *Config) error {
	rgCmd := exec.Command(
		"rg",
		"--hidden",
		"--glob", "!**/.git/**",
		"--line-number",
		"--no-heading",
		".",
		cfg.NotesDir,
	)

	fzfCmd := exec.Command(
		"fzf",
		"--delimiter", ":",
		"--with-nth", "1,2,3..",
		"--preview", "bat --style=numbers --color=always --line-range :500 {1} --highlight-line {2}",
	)

	rgOut, err := rgCmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create pipe: %w", err)
	}

	fzfCmd.Stdin = rgOut
	fzfOut, err := fzfCmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create pipe: %w", err)
	}

	if err := rgCmd.Start(); err != nil {
		return fmt.Errorf("failed to start ripgrep: %w", err)
	}

	if err := fzfCmd.Start(); err != nil {
		return fmt.Errorf("failed to start fzf: %w", err)
	}

	scanner := bufio.NewScanner(fzfOut)
	var selection string
	if scanner.Scan() {
		selection = scanner.Text()
	}

	rgCmd.Wait()
	fzfCmd.Wait()

	if selection == "" {
		return nil
	}

	parts := strings.SplitN(selection, ":", 3)
	if len(parts) < 2 {
		return fmt.Errorf("invalid selection format")
	}

	filepath := parts[0]
	lineNum := parts[1]

	cmd := exec.Command(cfg.Editor, fmt.Sprintf("+%s", lineNum), filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
