package renamer

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestFiles() (string, error) {

	dir, err := os.MkdirTemp("", "renamer_test")
	if err != nil {
		return "", err
	}

	// 创建一些测试文件
	files := []string{"test1.txt", "test2.txt"}
	for _, file := range files {
		filePath := filepath.Join(dir, file)
		if err := os.WriteFile(filePath, []byte("test"), 0666); err != nil {
			return dir, err
		}
	}
	return dir, nil
}

func TestRenameFiles(t *testing.T) {

	tests := []struct {
		renameType string
		newName    string
		prefix     string
		suffix     string
		wantFiles  []string
	}{
		{"full", "new", "", "", []string{"new-1.txt", "new-2.txt"}},
		{"prefix", "", "pre_", "", []string{"pre_test1.txt", "pre_test2.txt"}},
		{"suffix", "", "", "_suf", []string{"test1_suf.txt", "test2_suf.txt"}},
	}

	for _, test := range tests {
		t.Run(test.renameType, func(t *testing.T) {
			dir, err := setupTestFiles()
			if err != nil {
				t.Fatalf("Failed to setup test files: %v", err)
			}

			r := NewRenamer(dir)
			r.SetRenameType(test.renameType)
			r.SetNewName(test.newName)
			r.SetPrefix(test.prefix)
			r.SetSuffix(test.suffix)
			r.SelectAll = true
			if err := r.RenameFiles(); err != nil {
				t.Errorf("RenameFiles() failed: %v", err)
				return
			}

			for _, want := range test.wantFiles {
				if _, err := os.Stat(filepath.Join(dir, want)); os.IsNotExist(err) {
					t.Errorf("Expected file %s does not exist", want)
				}
			}
			defer os.RemoveAll(dir) // 测试完成后清理
		})
	}
}
