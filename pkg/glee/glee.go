package glee

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var ErrGitRepoNotFound = errors.New("git repository not found")
var ErrNoEditorAvailable = errors.New("unable to open exclude file with any available text editor")

func AddExcludes(entries []string) error {
	root, excludeFile, err := getGitRootAndExcludePath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(excludeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("opening exclude file: %w", err)
	}
	defer f.Close()

	for _, entry := range entries {
		if err := excludeEntry(entry, root, f); err != nil {
			return fmt.Errorf("excluding entry %q: %w", entry, err)
		}
	}

	return nil
}

func ListExcludes() error {
	_, excludeFile, err := getGitRootAndExcludePath()
	if err != nil {
		return err
	}

	f, err := os.Open(excludeFile)
	if err != nil {
		return fmt.Errorf("opening exclude file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			fmt.Println(line)
		}
	}

	return scanner.Err()
}

func ClearExcludes() error {
	_, excludeFile, err := getGitRootAndExcludePath()
	if err != nil {
		return err
	}

	if err := os.Truncate(excludeFile, 0); err != nil {
		return fmt.Errorf("truncating exclude file: %w", err)
	}

	return nil
}

func EditExcludes() error {
	_, excludeFile, err := getGitRootAndExcludePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(excludeFile); err != nil {
		return fmt.Errorf("accessing exclude file: %w", err)
	}

	for _, editor := range getEditorsList() {
		cmd := exec.Command(editor, excludeFile)
		if err := cmd.Start(); err == nil {
			return nil
		}
	}

	return ErrNoEditorAvailable
}

func excludeEntry(entry, root string, f *os.File) error {
	base, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current working directory: %w", err)
	}

	absolute := filepath.Join(base, entry)
	relative, err := filepath.Rel(root, absolute)
	if err != nil {
		return fmt.Errorf("building relative exclude path: %w", err)
	}

	if _, err := fmt.Fprintln(f, relative); err != nil {
		return fmt.Errorf("writing to exclude file: %w", err)
	}

	return nil
}

func getGitRootAndExcludePath() (string, string, error) {
	root, err := findClosestGitRoot()
	if err != nil {
		return "", "", fmt.Errorf("finding git root: %w", err)
	}
	excludePath := filepath.Join(root, ".git", "info", "exclude")
	return root, excludePath, nil
}

func findClosestGitRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getting current working directory: %w", err)
	}

	for dir != "/" {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}
		dir = filepath.Dir(dir)
	}

	return "", ErrGitRepoNotFound
}

func getEditorsList() []string {
	var editors []string

	// Check for $EDITOR environment variable
	if envEditor := os.Getenv("EDITOR"); envEditor != "" {
		editors = append(editors, envEditor)
	}

	// Add default system editors
	switch runtime.GOOS {
	case "darwin":
		editors = append(editors, "open", "TextEdit", "vim", "nano")
	case "windows":
		editors = append(editors, "notepad", "wordpad")
	default: // Linux and other Unix-like systems
		editors = append(editors, "xdg-open", "gedit", "kate", "kwrite", "nano", "vim", "emacs")
	}

	return editors
}
