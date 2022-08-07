package main

import (
	"github.com/aflores04/chat/src/websocket"
	"github.com/alecthomas/inject"
)

func Start(
	websocketServer websocket.WebsocketServer,
) {
	websocketServer.RunOnPort("8000")
}

func main() {
	injector := inject.New()
	injector.Install(
		&websocket.PoolModule{},
		&websocket.WebsocketServerModule{},
	)
	injector.Call(Start)
}
