package code

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var unitsByPower = map[float64]string{
	0:  "B",
	10: "KB",
	20: "MB",
	30: "GB",
	40: "TB",
	50: "PB",
	60: "EB",
}

func IsEntryVisible(entryName string) bool {
	return !strings.HasPrefix(entryName, ".")
}

func GetSizeString(byteSize int64, isHumanFormat bool) string {
	fSize := float64(byteSize)
	power := math.Log2(fSize)
	uPower := math.Floor(power*0.1) * 10
	if !isHumanFormat || uPower == 0 {
		return fmt.Sprintf("%d%s", byteSize, unitsByPower[uPower])
	}
	if int(uPower) < len(unitsByPower)*10 {
		fSize = math.Pow(2, power-uPower)
		return fmt.Sprintf("%.1f%s", fSize, unitsByPower[uPower])
	}
	return "Размер файла превышает 1024EB"
}

func GetDirElements(path string, includeHiddenElements bool) []os.DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return nil
	}
	if includeHiddenElements {
		return entries
	}
	visibleElements := []os.DirEntry{}
	for _, entry := range entries {
		if IsEntryVisible(entry.Name()) {
			visibleElements = append(visibleElements, entry)
		}
	}
	return visibleElements
}
func getFileSize(file os.DirEntry, includeHiddenElement bool) int64 {
	fileInfo, _ := file.Info()
	fileName := fileInfo.Name()
	excludeFile := !includeHiddenElement && !IsEntryVisible(fileName)
	if excludeFile {
		return 0
	}
	return fileInfo.Size()
}

func GetDirSize(path string, includeHiddenElements, recursive bool) int64 {
	var dirSize int64
	entries := GetDirElements(path, includeHiddenElements)
	for _, entry := range entries {
		if entry.IsDir() {
			if recursive {
				path = filepath.Join(path, entry.Name())
				dirSize += GetDirSize(path, includeHiddenElements, recursive)
			}
		} else {
			dirSize += getFileSize(entry, includeHiddenElements)
		}
	}
	return dirSize
}

func GetPathSize(path string, human, all, recursive bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	var pathSize int64

	if info.IsDir() {
		pathSize = GetDirSize(path, all, recursive)
	} else {
		pathSize = info.Size()
	}

	return GetSizeString(pathSize, human), nil
}
