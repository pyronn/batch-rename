package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Renamer holds the configuration for the renaming operation.
type Renamer struct {
	Directory     string
	RenameType    string
	NewName       string
	Prefix        string
	Suffix        string
	UseRegex      bool
	RegexPattern  string
	RenameExt     bool
	SelectAll     bool
	SelectedFiles []string
}

// NewRenamer creates a new Renamer instance with default values.
func NewRenamer(directory string) *Renamer {
	return &Renamer{
		Directory: directory,
	}
}

// SetRenameType sets the type of renaming operation.
func (r *Renamer) SetRenameType(renameType string) {
	r.RenameType = renameType
}

// SetNewName sets the new name for the 'full rename' operation.
func (r *Renamer) SetNewName(newName string) {
	r.NewName = newName
}

// SetPrefix sets the prefix for the 'prefix rename' operation.
func (r *Renamer) SetPrefix(prefix string) {
	r.Prefix = prefix
}

// SetSuffix sets the suffix for the 'suffix rename' operation.
func (r *Renamer) SetSuffix(suffix string) {
	r.Suffix = suffix
}

// EnableRegex enables the use of a regex pattern for file selection.
func (r *Renamer) EnableRegex(pattern string) {
	r.UseRegex = true
	r.RegexPattern = pattern
}

// EnableRenameExt enables renaming of the file extension.
func (r *Renamer) EnableRenameExt() {
	r.RenameExt = true
}

// RenameFiles performs the renaming operation based on the Renamer configuration.
func (r *Renamer) RenameFiles() (err error) {
	var fileNames []string

	var regex *regexp.Regexp
	if r.UseRegex {
		regex, err = regexp.Compile(r.RegexPattern)
		if err != nil {
			return fmt.Errorf("failed to compile regex pattern: %w", err)
		}
	}

	if r.SelectAll {
		files, err := os.ReadDir(r.Directory)
		if err != nil {
			return fmt.Errorf("failed to read directory: %w", err)
		}
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if r.UseRegex && !regex.MatchString(f.Name()) {
				continue // Skip files that don't match the regex
			}
			fileNames = append(fileNames, f.Name())
		}
	} else {
		fileNames = r.SelectedFiles
	}

	newFilenames, err := r.generateNewFileNames(fileNames)
	if err != nil {
		return err
	}

	for i, originalName := range fileNames {
		originalPath := filepath.Join(r.Directory, originalName)
		newPath := filepath.Join(r.Directory, newFilenames[i])

		if err := os.Rename(originalPath, newPath); err != nil {
			return fmt.Errorf("failed to rename file %s to %s: %w", originalName, newFilenames[i], err)
		}
	}

	return nil
}

func (r *Renamer) generateNewFileNames(files []string) ([]string, error) {

	var newFilenames []string

	count := 1
	for _, file := range files {

		originalName := file
		extension := filepath.Ext(originalName)
		baseName := strings.TrimSuffix(originalName, extension)

		newName := ""
		switch r.RenameType {
		case "full":
			baseNewName := r.NewName
			newExtension := "" // 用于存储可能从NewName解析出的新扩展名

			// 检查NewName是否包含扩展名
			if dotIndex := strings.LastIndex(r.NewName, "."); dotIndex != -1 && dotIndex != 0 {
				baseNewName = r.NewName[:dotIndex]  // NewName的基本部分（不含扩展名）
				newExtension = r.NewName[dotIndex:] // NewName的扩展名部分
			}

			// 如果计数大于0，我们需要在文件名中添加序号
			if count > 0 {
				baseNewName += "-" + strconv.Itoa(count)
			}

			// 根据RenameExt决定是否保留原始扩展名
			if !r.RenameExt && extension != "" && newExtension == "" {
				newExtension = extension
			}

			newName = baseNewName + newExtension

		case "prefix":
			newName = r.Prefix + originalName
		case "suffix":
			if r.RenameExt {
				newName = baseName + r.Suffix
			} else {
				newName = baseName + r.Suffix + extension
			}
		}
		count++
		newFilenames = append(newFilenames, newName)
	}

	return newFilenames, nil

}
