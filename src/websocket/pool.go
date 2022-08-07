package websocket

import "fmt"

type PoolModule struct {
	Handlers []MessageHandler
}

func (m *PoolModule) ProvidePool() *Pool {
	return &Pool{
		Handlers:   m.Handlers,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

type Pool struct {
	// Handlers to catch messages received from websocket
	Handlers []MessageHandler

	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			break
		case message := <-pool.Broadcast:
			// send messages to all handlers
			for _, handler := range pool.Handlers {
				handler.Handle(message)
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
