package main

import (
	"github.com/aflores04/chat/backend/src/chat_messages/db"
	"github.com/aflores04/chat/backend/src/chat_messages/service"
	"github.com/aflores04/chat/backend/src/chat_websocket_server/handler"
	"github.com/aflores04/chat/backend/src/mongodb"
	"github.com/aflores04/chat/backend/src/rabbitmq"
	"github.com/aflores04/chat/backend/src/websocket"
	"github.com/alecthomas/inject"
	"log"
)

func Start(
	websocketServer websocket.WebsocketServer,
	messageHandler handler.MessageHandler,
	stockCommandHandler handler.StockCommandHandler,
) {
	websocketServer.AddMessageHandler(messageHandler)
	websocketServer.AddMessageHandler(stockCommandHandler)
	log.Println("Chat Websocket server started ...")
	websocketServer.RunOnPort("8010")
}

func main() {
	injector := inject.New()
	injector.Install(
		&mongodb.MongoModule{},
		&rabbitmq.RabbitModule{},
		&db.ChatMessagesRepositoryModule{},
		&service.ChatMessagesServiceModule{},
		&websocket.PoolModule{},
		&handler.MessageHandlerModule{},
		&handler.StockCommandHandlerModule{},
		&websocket.WebsocketServerModule{},
	)
	injector.Call(Start)
}
