package code
import (
"os"
"fmt"
)
func GetPathSize(filePath string) (string, error) {
info, err := os.Lstat(filePath)
if err != nil {
return "", err
}
if info.IsDir() {
entries, _ :=os.ReadDir(filePath)
var dirSize int64 = 0
for _, entry := range entries {
entryInfo, _ := entry.Info()
dirSize += entryInfo.Size()
}
return fmt.Sprintf("%dB", dirSize), nil
} else {
return fmt.Sprintf("%dB", info.Size()), nil
}
}
