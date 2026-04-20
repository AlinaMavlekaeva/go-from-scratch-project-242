package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ConvertSize(size int64, convert bool) string {
	units := map[int]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
		6: "EB",
	}
	if !convert || size < 1024 {
		return fmt.Sprintf("%d%s", size, units[0])
	}
	size = size / 1024
	for i := 1; i < 6; i++ {
		if size < 1024 {
			return fmt.Sprintf("%.1f%s", float64(size), units[i])
		}
		size = size / 1024
	}
	return fmt.Sprintf("%.1f%s", float64(size), units[6])
}
func GetDirElements(path string, all bool) []os.DirEntry {
	selected := []os.DirEntry{}
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return nil
	}
	if all {
		return entries
	}
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") {
			selected = append(selected, entry)
		}
	}
	return selected
}
func GetDirSize(path string, recursive, all bool) int64 {
	var dirSize int64 = 0
	entries := GetDirElements(path, all)
	for _, entry := range entries {
		if entry.IsDir() {
			if recursive {
				path = filepath.Join(path, entry.Name())
				dirSize += GetDirSize(path, recursive, all)
			}
		} else {
			fileInfo, _ := entry.Info()
			dirSize += fileInfo.Size()
		}
	}
	return dirSize
}
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return ConvertSize(info.Size(), human), nil
	}
	dirSize := GetDirSize(path, recursive, all)
	return ConvertSize(dirSize, human), nil

}
