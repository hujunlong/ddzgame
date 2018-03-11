package internal

import (
 	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/game"
	"fmt"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.Hello{},handleHello)
}

func handleHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)
	
	log.Debug("hello %v", m.Name)
	fmt.Println("recive hello:", m)
	
	//给game发送消息
	game.ChanRPC.Go("NewAgent", a)
	
	result ,status := game.ChanRPC.Call1("CloseAgent",a)
	fmt.Println("result:",result,"status:",status)
	
	a.WriteMsg(&msg.Hello{
		Name:"client",
	})
}