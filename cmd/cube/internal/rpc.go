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
				Name:    "methods",
				Aliases: []string{"m"},
				Usage:   "指定生成的rpc方法名称，逗号分隔，支持多个；会生成该接口的rpc声明以及message声明",
			},
			&cli.StringFlag{
				Name:    "request",
				Aliases: []string{"req"},
				Usage:   "rpc方法入参后缀，默认Request",
			},
			&cli.StringFlag{
				Name:    "reply",
				Aliases: []string{"rep"},
				Usage:   "rpc方法返回值后缀，默认Reply",
			},
		},
		Action: action.NewRPCAction().GenerateRpcFile,
	}
}
