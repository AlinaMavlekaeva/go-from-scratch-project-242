package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   true,
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fileName := cmd.Args().Get(0)
			fileSize, err := code.GetPathSize(fileName, cmd.Bool("human"), cmd.Bool("all"), cmd.Bool("recursive"))
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
