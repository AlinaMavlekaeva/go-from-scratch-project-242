package code
import (
"os"
"fmt"
)
func GetPathSize(path, string) (string) {
info, _ := os.Lstat(path)
if !info.IsDir {
        return fmt.Sprintf("%dB", info.Size())
}
else {
entries, _ :=os.ReadDir(path)
dirSize := 0
for _, entry := range entries {
entryInfo, _ := entry.Info()
dirSize += entryInfo.Size()
}
return fmt.Sprintf("%dB", dirSize)
}
