package action

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

type rpcAction struct{}

func NewRPCAction() *rpcAction {
	return &rpcAction{}
}

var serviceFmt = "service %s {\n%s\n}\n"
var rpcFmt = "  rpc %s (%s) returns (%s){\n    option (google.api.http) = {\n    };\n  };\n"
var msgFmt = "message %s{\n\n}\n\n"

const DefaultRequest = "Request"
const DefaultReply = "Reply"
const DefaultSvc = "DefaultService"

func (a *rpcAction) GenerateRpcFile(c *cli.Context) error {
	methods := c.String("methods")
	if methods == "" {
		methods = c.String("m")
	}
	if methods == "" {
		fmt.Fprintf(os.Stderr, "请指定rpc子命令 -m参数\n")
		return errors.New("可参考cube help的对应子命令 cube rpc help")
	}

	svc := c.String("service")
	if svc == "" {
		svc = c.String("s")
	}
	if svc == "" {
		svc = DefaultSvc
	}
	replySuffix := DefaultReply
	requestSuffix := DefaultRequest

	if req := c.String("request"); req != "" {
		requestSuffix = req
	} else if req = c.String("req"); req != "" {
		requestSuffix = req
	}
	if resp := c.String("reply"); resp != "" {
		replySuffix = resp
	} else if resp = c.String("rep"); resp != "" {
		requestSuffix = resp
	}
	resultRpc := ""
	messageRes := ""
	methodList := strings.Split(methods, ",")
	for _, val := range methodList {
		if val == "" {
			continue
		}
		req := val + requestSuffix
		reply := val + replySuffix
		resultRpc += fmt.Sprintf(rpcFmt, val, req, reply)
		messageRes += fmt.Sprintf(msgFmt, req)
		messageRes += fmt.Sprintf(msgFmt, reply)
	}
	serviceRes := fmt.Sprintf(serviceFmt, svc, resultRpc)
	fmt.Fprintf(os.Stdout, serviceRes)
	fmt.Fprintf(os.Stdout, "\n")
	fmt.Fprintf(os.Stdout, messageRes)
	return nil

}
