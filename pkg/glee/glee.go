package glee

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func AddExcludes(entries []string) error {
	root, err := findClosestGitRoot()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if err := excludeEntry(entry, root); err != nil {
			return err
		}
	}

	return nil
}

func ListExcludes() error {
	root, err := findClosestGitRoot()
	if err != nil {
		return err
	}

	f, err := os.Open(getRootExcludeFile(root))
	if err != nil {
		return fmt.Errorf("cannot open exclude file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 && !strings.HasPrefix(line, "#") {
			fmt.Println(line)
		}
	}

	return scanner.Err()
}

func ClearExcludes() error {
	root, err := findClosestGitRoot()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(getRootExcludeFile(root), os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("cannot truncate exclude file: %w", err)
	}
	defer f.Close()

	return nil
}

func EditExcludes() error {
	root, err := findClosestGitRoot()
	if err != nil {
		return err
	}

	excludeFile := getRootExcludeFile(root)
	_, err = os.Stat(excludeFile)
	if err != nil {
		return fmt.Errorf("cannot access exclude file: %w", err)
	}

	for _, editor := range getEditorsList() {
		cmd := exec.Command(editor, excludeFile)
		err := cmd.Start()
		if err == nil {
			return nil // Successfully opened the file
		}
	}

	return errors.New("unable to open exclude file with any available text editor")
}

func excludeEntry(entry, root string) error {
	base, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get current working directory: %w", err)
	}

	absolute := path.Join(base, entry)
	relative, err := filepath.Rel(root, absolute)
	if err != nil {
		return fmt.Errorf("cannot build a relative exclude path: %w", err)
	}

	f, err := os.OpenFile(getRootExcludeFile(root), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open exclude file: %w", err)
	}
	defer f.Close()

	if _, err = f.Write([]byte(relative + "\n")); err != nil {
		return fmt.Errorf("cannot write to exclude file: %w", err)
	}

	return nil
}

func findClosestGitRoot() (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cannot get current working directory: %w", err)
	}

	for current != "/" {
		candidate := path.Join(current, ".git")

		if stat, err := os.Stat(candidate); err == nil && stat.IsDir() {
			return current, nil
		}

		current = filepath.Dir(current)
	}

	return "", errors.New("git repository not found")
}

func getRootExcludeFile(root string) string {
	return path.Join(root, ".git", "info", "exclude")
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
