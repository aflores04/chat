package websocket

// MessageHandler handler for messages received from websocket
// use this interface to create your own handlers
type MessageHandler interface {
	Handle(message WebsocketMessage)
}
