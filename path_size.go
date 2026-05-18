package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isVisible(entryName string) bool {
	return !strings.HasPrefix(entryName, ".")
}

func formatSize(byteSize int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", byteSize)
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	val, i := float64(byteSize), 0
	for val >= 1024 && i < len(units)-1 {
		val /= 1024
		i++
	}
	if i == 0 {
		return fmt.Sprintf("%d%s", byteSize, units[i])
	}
	return fmt.Sprintf("%.1f%s", val, units[i])
}
func getDirSize(path string, all, recursive bool) (int64, error) {
	var dirSize int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, entry := range entries {
		if all || isVisible(entry.Name()) {
			sub := filepath.Join(path, entry.Name())
			if entry.IsDir() {
				if recursive {
					subSize, err := getDirSize(sub, all, recursive)
					if err != nil {
						return 0, err
					}
					dirSize += subSize
				}
			} else {
				info, err := os.Stat(sub)
				if err != nil {
					return 0, err
				}
				dirSize += info.Size()
			}
		}
	}
	return dirSize, nil
}
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if !all && !isVisible(filepath.Base(path)) {
		return "0", nil
	}
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		dirSize, err := getDirSize(path, all, recursive)
		if err != nil {
			return "", err
		}
		return formatSize(dirSize, human), nil
	}
	return formatSize(info.Size(), human), nil
}
