package WebSockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"context"
	"RealTimeChatApp/Backend/Redis"
)

var ConnectionUpgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(*http.Request) bool{
		return true
	},
}

type Handler struct {
	Hub *Hub
}

func CreateHandler(Hub *Hub) *Handler{
	return &Handler{
		Hub: Hub,
	}
}

func (Handler *Handler) HandleWebSocketConnection(GinContext *gin.Context){
	SocketConnection, err := ConnectionUpgrader.Upgrade(GinContext.Writer, GinContext.Request, nil)
	if err != nil {
		log.Println("Connection upgrade error: ", err)
		return
	}

	log.Println("Upgraded to websocket!")

	ContextValue, Exists := GinContext.Get("UserID")

	if !Exists {
		GinContext.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to perform this action"})
	}

	UserID := ContextValue.(bson.ObjectID)
	Client := &Client{
		UserID: UserID,
		Connection: SocketConnection,
		SendChannel: make(chan []byte),
	}

	Handler.Hub.RegisterClient(Client)

	sub := Redis.Client.Subscribe(context.TODO(), "Message")

	go WritePump(Client)
	go ReadPump(Handler.Hub)
	for {
		msg, err := sub.ReceiveMessage(context.TODO())
		if err != nil {
			log.Println(err)
		}

		log.Println(msg)

	// TODO: Find a way to use the hub to find the correct active connection and send the message there
	}
}

func WritePump(Client *Client){
	// TODO: Implement the write pump
	// Its purpose is to utilize the socket channel to write data to from the sendchannel of the target client

}

func ReadPump(Hub *Hub){
	// TODO: Implement the Read Pump
	// Its purpose is to take notice of any user changes, i.e., disconnected, is typing, etc.
}
