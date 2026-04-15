package main
import (
"os"
"context"
"github.com/urfave/cli/v3"
"fmt"
"code"
)
func main() {
cmd := &cli.Command {
Flags: []cli.Flag {
&cli.BoolFlag {
 Name: "human",
 Aliases: []string{"H"},
Value: true,
 Usage: "human-readable sizes (auto-select unit)",
},
},
Action: func(ctx context.Context, cmd *cli.Command) error {
fileName := cmd.Args().Get(0)
fileSize, err := code.GetPathSize(fileName, cmd.Bool("human"))
if err != nil {
return err
}
result := fmt.Sprintf("%s\t%s", fileSize, fileName)
fmt.Println(result)
return nil
},
}
if err := cmd.Run(context.Background(), os.Args); err != nil {
fmt.Println("Ошибка: ", err)
}
}
