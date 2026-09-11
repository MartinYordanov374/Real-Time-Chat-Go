package WebSockets

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/gorilla/websocket"
	"sync"
)
type Hub struct {
	ActiveClients map[*Client]bool
	Mutex sync.RWMutex
	broadcast chan []byte
}

// TODO: Make the channel buffered
type Client struct {
	UserID bson.ObjectID
	Connection *websocket.Conn
	SendChannel chan []byte
}
