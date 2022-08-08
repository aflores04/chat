package handler

import (
	"context"
	"encoding/json"
	"github.com/AlekSi/pointer"
	"github.com/aflores04/chat/backend/src/rabbitmq"
	"github.com/aflores04/chat/backend/src/websocket"
	"log"
	"regexp"
	"strings"
)

const (
	StockCommandsQueue       = "stock-commands"
	StockCommandPrefix       = "/stock"
	StockCommandDivider      = "="
	StockCommandCodePosition = 1
)

type StockCommandHandlerModule struct{}

func (*MessageHandlerModule) ProvideCommandHandler(wsServer websocket.WebsocketServer, rabbit rabbitmq.RabbitQueue) StockCommandHandler {
	return &stockCommandHandler{
		rabbit:   rabbit,
		wsServer: wsServer,
	}
}

type StockCommandHandler interface {
	Handle(message websocket.WebsocketMessage)
}

type stockCommandHandler struct {
	rabbit   rabbitmq.RabbitQueue
	wsServer websocket.WebsocketServer
}

func (h stockCommandHandler) Handle(wsMessage websocket.WebsocketMessage) {
	ctx := context.Background()

	if !IsStockCommand(*wsMessage.Payload.Body) {
		return
	}

	// get stock code from command
	// add stock code to body in payload to be treated by the bot
	stockCode := strings.Split(*wsMessage.Payload.Body, StockCommandDivider)
	wsMessage.Payload.Body = pointer.ToString(stockCode[StockCommandCodePosition])

	b, err := json.Marshal(wsMessage.Payload)
	if err != nil {
		log.Println("error marshal websocket message: ", err)
	}

	queue := h.rabbit.CreateQueue(StockCommandsQueue)
	h.rabbit.Publish(ctx, queue, string(b))
}

func IsStockCommand(bodyMessage string) bool {
	match, err := regexp.MatchString(StockCommandPrefix, bodyMessage)
	if err != nil {
		log.Println(err)
		return false
	}

	if !match {
		return false
	}

	return true
}
