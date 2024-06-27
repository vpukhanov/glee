package glee

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
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

	for current != "." {
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
