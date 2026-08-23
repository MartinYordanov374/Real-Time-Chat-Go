package WebSockets

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"github.com/gorilla/websocket"
	"sync"
)
type Hub struct {
	ActiveClients map[*Client]bool
	Mutex sync.RWMutex
}

type Client struct {
	UserID bson.ObjectID
	Connection *websocket.Conn
	SendChannel chan []byte
}
