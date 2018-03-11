package internal

import (
	"github.com/name5566/leaf/gate"
	"fmt"
	"time"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	fmt.Println("---rpcNewAgent---")
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) interface{}{
	fmt.Println("---rpcCloseAgent---")
	time.Sleep(300*time.Millisecond)
	a := args[0].(gate.Agent)
	_ = a
	return 1
}