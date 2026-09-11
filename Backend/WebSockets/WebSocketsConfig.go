package WebSockets

import (
	"RealTimeChatApp/Backend/GlobalVariables"
	"RealTimeChatApp/Backend/Redis"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var ConnectionUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

type Handler struct {
	Hub *Hub
}

func CreateHandler(Hub *Hub) *Handler {
	return &Handler{
		Hub: Hub,
	}
}

func (Handler *Handler) HandleWebSocketConnection(GinContext *gin.Context) {
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
		UserID:      UserID,
		Connection:  SocketConnection,
		SendChannel: make(chan []byte),
	}

	Handler.Hub.RegisterClient(Client)
	// TODO: Move the subscription out of the socket connection handler to avoid duplications
	sub := Redis.Client.Subscribe(context.Background(), "Message")

	go Client.WritePump()
	go Client.ReadPump()
	for {
		msg, err := sub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		var PayloadData GlobalVariables.RedisMessage
		err = json.Unmarshal([]byte(msg.Payload), &PayloadData)
		if err != nil {
			log.Println(err)
		}
		// TODO: Consider what happens when the user is offline when they are sent a message
		// Store the messages in the DB, cache the last 100 messages in Redis
		// When they come back check if any new messages compared to the latest Redis cached one have arrived
		// If yes, fetch directly from the DB, otherwise fetch from Redis
		//
		// TODO: Implement sent/delievered functionality
		// TODO: Implement a writing indicator functionality
		// TODO: Implement an online/offline status indicator functionality
		SendMessageToClient(Handler.Hub, PayloadData.SenderID, PayloadData.ReceiverID, PayloadData.Content)
	}
}

func (Client *Client) WritePump() {
	for {
		msg := <-Client.SendChannel
		err := Client.Connection.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("An error occurred")
			log.Println(err)
		}
		// TODO: Handle what happens after closing the channels
	}
}

func (Client *Client) ReadPump() {
	// TODO: Implement the Read Pump
	// Its purpose is to take notice of any user changes, i.e., disconnected, is typing, etc.
}

func SendMessageToClient(Hub *Hub, SenderID bson.ObjectID, ReceiverID bson.ObjectID, Message string) {
	// TODO: Find a way to make this O(1)
	for client := range Hub.ActiveClients {
		if client.UserID == SenderID || client.UserID == ReceiverID {
			client.SendChannel <- []byte(Message)
		}
	}
}
