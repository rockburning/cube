package action

import (
	"github.com/urfave/cli/v2"
	"testing"
)

func Test_rpcAction_GenerateRpcFile(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &rpcAction{}
			if err := a.GenerateRpcFile(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GenerateRpcFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
