package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aflores04/chat/src/rabbitmq"
	"github.com/aflores04/chat/src/stock"
	ws "github.com/aflores04/chat/src/websocket"
	"github.com/aflores04/chat/src/websocket_server/handler"
	"github.com/alecthomas/inject"
	"github.com/alecthomas/kingpin"
	"github.com/gorilla/websocket"
	"log"
)

const (
	CommandBotUsername = "bot"
)

type Application struct {
	websocketServerUrl string
	rabbitQueueName    string
}

func (a *Application) Configure() {
	kingpin.Flag("websocket-server", "Websocket server to listen").
		Default("ws://localhost:8010").
		Envar("WEBSOCKET_SERVER").
		StringVar(&a.websocketServerUrl)
	kingpin.Flag("rabbit-queue", "Rabbit queue to listen").
		Default("stock-commands").
		Envar("RABBIT_QUEUE_NAME").
		StringVar(&a.rabbitQueueName)
	kingpin.Parse()
}

func (a *Application) Start(
	rabbit rabbitmq.RabbitQueue,
	stockService stock.StockService,
) {
	queue := rabbit.CreateQueue(a.rabbitQueueName)

	// Dial to websocket server
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(a.websocketServerUrl, nil)
	if err != nil {
		log.Fatal("error dialing to websocket server: ", err.Error())
	}

	log.Println("Starting bot service...")
	log.Printf("Listening %s queue", handler.StockCommandsQueue)
	quit := make(chan bool)
	go func() {
		for event := range rabbit.PollMessages(queue) {
			var message []byte

			// Parse received command
			receivedPayload := &ws.Payload{}
			_ = json.NewDecoder(bytes.NewReader(event.Body)).Decode(&receivedPayload)

			log.Println("Command received in bot: ", receivedPayload)

			// get stock data fro service
			stockData, err := stockService.GetStockByCode(receivedPayload.Body)
			if err != nil {
				log.Println("error getting stock data")
				message = createMessage("stock not found")
			} else {
				// create byte message to be sent
				body := fmt.Sprintf("%s quote is %s per share.",
					stockData[stock.SymbolPosition],
					stockData[stock.OpenPosition],
				)
				message = createMessage(body)
			}

			// send message to websocket
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("error writing message: ", err)
				return
			}
		}
		quit <- true
	}()

	<-quit
	log.Println("bot application terminated")
}

func createMessage(body string) []byte {
	sentPayload := &ws.Payload{
		Username: CommandBotUsername,
		Body:     body,
	}
	b, err := json.Marshal(&sentPayload)
	if err != nil {
		log.Println("error marshall payload from bot: ", err)
	}

	return b
}

func main() {
	app := &Application{}
	app.Configure()

	injector := inject.New()
	injector.Install(
		&rabbitmq.RabbitModule{},
		&stock.StockServiceModule{},
	)
	injector.Call(app.Start)
}
