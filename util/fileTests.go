package util

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"regexp"
)

func requireFileOrDir(rootDir string, items []string, needsDir bool) bool {
	for _, checkPath := range items {
		fullPath := path.Join(rootDir, checkPath)

		if stat, err := os.Stat(fullPath); err != nil || (needsDir && !stat.IsDir()) || (!needsDir && stat.IsDir()) {
			itemType := "file"
			if needsDir {
				itemType = "directory"
			}
			fmt.Printf("required %s missing: %s\n", itemType, checkPath)
			return false
		}
	}

	return true
}

func requireRegexFileOrDirMatch(rootDir string, namePattern string, needsDir bool) (bool, string) {
	entries, err := os.ReadDir(rootDir)

	if err != nil {
		slog.Error(fmt.Sprintf("Error occurred when reading directory '%s': %s", rootDir, err))
		return false, ""
	}

	for _, entry := range entries {
		match, _ := regexp.MatchString(namePattern, entry.Name())

		if match && ((needsDir && entry.IsDir()) || (!needsDir && !entry.IsDir())) {
			return true, path.Join(rootDir, entry.Name())
		}
	}

	return false, ""
}

func RequireDirs(rootDir string, dirs []string) bool {
	return requireFileOrDir(rootDir, dirs, true)
}

func RequireFiles(rootDir string, files []string) bool {
	return requireFileOrDir(rootDir, files, false)
}

func RequireRegexDirMatch(rootDir string, namePattern string) (bool, string) {
	return requireRegexFileOrDirMatch(rootDir, namePattern, true)
}

func RequireRegexFileMatch(rootDir string, namePattern string) (bool, string) {
	return requireRegexFileOrDirMatch(rootDir, namePattern, false)
}