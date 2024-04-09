package main

import (
	"flag"
	"fmt"
	"github.com/pyronn/batch-rename/renamer"
	"os"
	"path/filepath"
)

func main() {
	// Define command-line flags
	directory := flag.String("dir", ".", "Directory containing files to rename")
	renameType := flag.String("type", "full", "Type of rename operation: full, prefix, suffix")
	newName := flag.String("name", "", "New name for 'full' rename operation")
	prefix := flag.String("prefix", "", "Prefix for 'prefix' rename operation")
	suffix := flag.String("suffix", "", "Suffix for 'suffix' rename operation")
	useRegex := flag.Bool("regex", false, "Use regular expression for file selection")
	regexPattern := flag.String("pattern", "", "Regular expression pattern for file selection")
	renameExt := flag.Bool("ext", false, "Rename file extension as well")
	selectAll := flag.Bool("all", false, "Select all files in the directory")

	flag.Parse()

	absPath, err := filepath.Abs(*directory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Create a new Renamer instance
	renamerInstance := renamer.NewRenamer(absPath)
	renamerInstance.SetRenameType(*renameType)
	renamerInstance.SetNewName(*newName)
	renamerInstance.SetPrefix(*prefix)
	renamerInstance.SetSuffix(*suffix)
	if *useRegex {
		renamerInstance.EnableRegex(*regexPattern)
	}
	if *renameExt {
		renamerInstance.EnableRenameExt()
	}

	// 如果设置了全选，那么将忽略正则表达式的设置，因为全选意味着操作目录下的所有文件
	if *selectAll {
		renamerInstance.UseRegex = false
		renamerInstance.SelectAll = true
	}

	// Execute the rename operation
	err = renamerInstance.RenameFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Files renamed successfully.")
}
