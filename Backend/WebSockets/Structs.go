package WebSockets

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/gorilla/websocket"

)
type Hub struct {
	ActiveClients map[*Client]bool
	RegisterClient chan *Client
	UnregisterClient chan *Client
}

type Client struct {
	UserID bson.ObjectID
	Connection *websocket.Conn
	SendChannel chan []byte
}
