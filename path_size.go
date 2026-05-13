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

func filterDirEntries(entries []os.DirEntry) []os.DirEntry {
	filtred := []os.DirEntry{}
	for _, entry := range entries {
		if isVisible(entry.Name()) {
			filtred = append(filtred, entry)
		}
	}
	return filtred
}
func getDirSize(path string, all, recursive bool) (int64, error) {
	var dirSize int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	if !all {
		entries = filterDirEntries(entries)
	}
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}
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
			if info.Mode()&os.ModeSymlink != 0 {
				symInfo, err := os.Stat(entry.Name())
				if err != nil {
					return 0, err
				}
				dirSize += symInfo.Size()
			}
			fileSize := info.Size()
			dirSize += fileSize
		}
	}
	return dirSize, nil
}
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	if !all && !isVisible(filepath.Base(path)) {
		return formatSize(0, human), nil
	}
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	var pathSize int64
	if info.IsDir() {
		dirSize, err := getDirSize(path, all, recursive)
		if err != nil {
			return "", err
		}
		pathSize = dirSize
	} else {
		pathSize = info.Size()
	}
	return formatSize(pathSize, human), nil
}
