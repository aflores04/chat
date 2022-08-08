package websocket

import (
	"fmt"
	"log"
)

type PoolModule struct {
	MessageHandlers []MessageHandler
}

func (m *PoolModule) ProvidePool() *Pool {
	return &Pool{
		MessageHandlers: m.MessageHandlers,
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		Clients:         make(map[*Client]bool),
		Broadcast:       make(chan WebsocketMessage),
	}
}

type Pool struct {
	// Handlers to catch messages received from websocket
	MessageHandlers []MessageHandler

	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan WebsocketMessage
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("{ Websocket } Client connected pool size: ", len(pool.Clients))
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("{ Websocket } Client disconnected pool size: ", len(pool.Clients))
			break
		case message := <-pool.Broadcast:
			log.Println("{ Websocket } Message received, sending to all clients: ", message)

			// send messages to all handlers
			for _, handler := range pool.MessageHandlers {
				go handler.Handle(message)
			}

			// send messages to all websocket clients
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
