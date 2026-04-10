package main
import (
"os"
"context"
"github.com/urfave/cli/v3"
"fmt"
"code"
)
func main() {
cmd := &cli.Command{
Name:   "hexlet-path-size"
Usage:  "Get path's size"
Action: func(ctx context.Context, cmd *cli.Command) error {
fileName := cmd.Args().Get(0)
fileSize := code.GetPathSize(fileName)
result := fmt.Sprintf("%s\t%s", fileSize, fileName)
fmt.Println(result)
return nil
        },
}
if err := cmd.Run(context.Background(), os.Args); err != nil {
log.Fatal(err)
        }
}
