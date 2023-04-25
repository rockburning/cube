package main

import (
	"fmt"
	"github.com/rockburning/cube/cmd/cube/internal"

	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "cube",
		Usage: "A simple tool for coding usage",
		Commands: []*cli.Command{
			internal.NewRpcCMD(),
		}}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
