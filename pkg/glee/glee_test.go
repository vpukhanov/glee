package glee

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// setupTestRepo creates a temporary directory with a .git folder
func setupTestRepo(t *testing.T) (string, func()) {
	t.Helper()
	tempDir, err := os.MkdirTemp("", "glee-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	gitDir := filepath.Join(tempDir, ".git", "info")
	if err := os.MkdirAll(gitDir, 0755); err != nil {
		t.Fatalf("Failed to create .git directory: %v", err)
	}

	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	return tempDir, cleanup
}

func TestAddExcludes(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change to the temp directory
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	entries := []string{"file1.txt", "dir/file2.txt"}
	if err := AddExcludes(entries); err != nil {
		t.Fatalf("AddExcludes failed: %v", err)
	}

	// Check if entries were added to the exclude file
	excludeFile := filepath.Join(tempDir, ".git", "info", "exclude")
	content, err := os.ReadFile(excludeFile)
	if err != nil {
		t.Fatalf("Failed to read exclude file: %v", err)
	}

	expectedContent := "file1.txt\ndir/file2.txt\n"
	if string(content) != expectedContent {
		t.Errorf("Exclude file content mismatch. Expected:\n%s\nGot:\n%s", expectedContent, string(content))
	}
}

func TestListExcludes(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change to the temp directory
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	// Add some excludes
	excludeFile := filepath.Join(tempDir, ".git", "info", "exclude")
	content := "file1.txt\n# Comment\ndir/file2.txt\n"
	if err := os.WriteFile(excludeFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write exclude file: %v", err)
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	if err := ListExcludes(); err != nil {
		t.Fatalf("ListExcludes failed: %v", err)
	}

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	expected := "file1.txt\ndir/file2.txt\n"
	if buf.String() != expected {
		t.Errorf("ListExcludes output mismatch. Expected:\n%s\nGot:\n%s", expected, buf.String())
	}
}

func TestClearExcludes(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Change to the temp directory
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	// Add some excludes
	excludeFile := filepath.Join(tempDir, ".git", "info", "exclude")
	content := "file1.txt\ndir/file2.txt\n"
	if err := os.WriteFile(excludeFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write exclude file: %v", err)
	}

	if err := ClearExcludes(); err != nil {
		t.Fatalf("ClearExcludes failed: %v", err)
	}

	// Check if file is empty
	newContent, err := os.ReadFile(excludeFile)
	if err != nil {
		t.Fatalf("Failed to read exclude file: %v", err)
	}

	if len(newContent) != 0 {
		t.Errorf("Exclude file not empty after clearing. Content: %s", string(newContent))
	}
}

func TestFindClosestGitRoot(t *testing.T) {
	tempDir, cleanup := setupTestRepo(t)
	defer cleanup()

	// Resolve symlinks in tempDir
	realTempDir, err := filepath.EvalSymlinks(tempDir)
	if err != nil {
		t.Fatalf("Failed to resolve symlinks: %v", err)
	}

	// Change to a subdirectory of the temp directory
	subDir := filepath.Join(realTempDir, "subdir", "anotherdir")
	os.MkdirAll(subDir, 0755)
	oldWd, _ := os.Getwd()
	os.Chdir(subDir)
	defer os.Chdir(oldWd)

	root, err := findClosestGitRoot()
	if err != nil {
		t.Fatalf("findClosestGitRoot failed: %v", err)
	}

	if root != realTempDir {
		t.Errorf("findClosestGitRoot returned incorrect root. Expected: %s, Got: %s", realTempDir, root)
	}
}

func TestGetEditorsList(t *testing.T) {
	oldEditor := os.Getenv("EDITOR")
	os.Setenv("EDITOR", "myeditor")
	defer os.Setenv("EDITOR", oldEditor)

	editors := getEditorsList()

	if editors[0] != "myeditor" {
		t.Errorf("First editor should be from EDITOR env var. Expected: myeditor, Got: %s", editors[0])
	}

	// Check if system editors are included (this will vary by OS)
	found := false
	for _, editor := range editors[1:] {
		if editor == "vim" || editor == "notepad" || editor == "gedit" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("System editors not found in the list: %v", editors)
	}
}
