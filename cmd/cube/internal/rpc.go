package internal

import (
	"github.com/rockburning/cube/pkg/action"

	"github.com/urfave/cli/v2"
)

func NewRpcCMD() *cli.Command {
	return &cli.Command{
		Name:  "rpc",
		Usage: "to generate rpc interface and message",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "service",
				Aliases: []string{"s"},
				Usage:   "指定生成的service名称，可忽略;默认为DefaultService",
			},
			&cli.StringFlag{
				Name:    "interface",
				Aliases: []string{"i"},
				Usage:   "指定生成的rpc名称，逗号分隔；会生成该接口的rpc声明以及message声明",
			},
		},
		Action: action.NewRPCAction().GenerateRpcFile,
	}
}
