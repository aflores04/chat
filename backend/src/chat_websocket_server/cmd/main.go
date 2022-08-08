package main

import (
	"github.com/aflores04/chat/src/chat/db"
	"github.com/aflores04/chat/src/chat/service"
	"github.com/aflores04/chat/src/mongodb"
	"github.com/aflores04/chat/src/rabbitmq"
	"github.com/aflores04/chat/src/websocket"
	"github.com/aflores04/chat/src/websocket_server/handler"
	"github.com/alecthomas/inject"
)

func Start(
	websocketServer websocket.WebsocketServer,
	messageHandler handler.MessageHandler,
	stockCommandHandler handler.StockCommandHandler,
) {
	websocketServer.AddMessageHandler(messageHandler)
	websocketServer.AddMessageHandler(stockCommandHandler)
	websocketServer.RunOnPort("8010")
}

func main() {
	injector := inject.New()
	injector.Install(
		&mongodb.MongoModule{},
		&rabbitmq.RabbitModule{},
		&db.ChatRepositoryModule{},
		&service.ChatServiceModule{},
		&websocket.PoolModule{},
		&handler.MessageHandlerModule{},
		&handler.StockCommandHandlerModule{},
		&websocket.WebsocketServerModule{},
	)
	injector.Call(Start)
}