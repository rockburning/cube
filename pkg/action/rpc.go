package action

import (
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

const Request = "Request"
const Reply = "Reply"

func (a *rpcAction) GenerateRpcFile(c *cli.Context) error {
	iface := c.String("interface")
	svc := c.String("service")
	if iface == "" {
		iface = c.String("i")
	}
	if svc == "" {
		svc = c.String("s")
	}
	if svc == "" {
		svc = "DefaultService"
	}
	resultRpc := ""
	messageRes := ""
	ifaceList := strings.Split(iface, ",")
	for _, val := range ifaceList {
		req := val + Request
		reply := val + Reply
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
