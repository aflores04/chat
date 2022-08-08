package websocket

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebsocketServerModule struct{}

func (*WebsocketServerModule) ProvideWebsocketServer(pool *Pool) WebsocketServer {
	return &websocketServer{
		pool: pool,
	}
}

type WebsocketServer interface {
	// RunOnPort start websocket server on specific port
	RunOnPort(port string)

	// AddHandler add handler for messages received from websocket
	AddMessageHandler(handler MessageHandler)
}

type websocketServer struct {
	pool *Pool
}

func (s *websocketServer) AddMessageHandler(handler MessageHandler) {
	s.pool.MessageHandlers = append(s.pool.MessageHandlers, handler)
}

func (s websocketServer) RunOnPort(port string) {
	r := chi.NewRouter()

	go s.pool.Start()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Fprintf(w, "%+v\n", err)
		}

		client := &Client{
			Conn: conn,
			Pool: s.pool,
		}

		s.pool.Register <- client
		client.Read()
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
