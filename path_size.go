package code
import (
"os"
"fmt"
)
func ConvertSize(size int64) string {
units := map[int]string{
0: "B",
1: "KB",
2: "MB",
3: "GB",
4: "TB",
5: "PB",
6: "EB",
}
if size < 1024 {
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
func GetPathSize(path string, human bool) (string, error) {
info, err := os.Lstat(path)
if err != nil {
return "", err
}
if info.IsDir() {
entries, _ :=os.ReadDir(path)
var dirSize int64 = 0
for _, entry := range entries {
entryInfo, _ := entry.Info()
dirSize += entryInfo.Size()
}
if human {
return ConvertSize(dirSize), nil
} else {
return fmt.Sprintf("%dB", dirSize), nil
}
}
if human {
return ConvertSize(info.Size()), nil
}
return fmt.Sprintf("%dB", info.Size()), nil
}
