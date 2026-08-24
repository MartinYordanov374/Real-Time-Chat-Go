package WebSockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"context"
	"RealTimeChatApp/Backend/Redis"
	"RealTimeChatApp/Backend/GlobalVariables"
	"encoding/json"
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

	go Client.WritePump()
	go Client.ReadPump()
	for {
		msg, err := sub.ReceiveMessage(context.TODO())
		if err != nil {
			log.Println(err)
		}

		var PayloadData GlobalVariables.RedisMessage
		err = json.Unmarshal([]byte(msg.Payload), &PayloadData)
		if err != nil{
			log.Println(err)
		}

		log.Println(PayloadData.UserID)
		SendMessageToClient(Handler.Hub, PayloadData.UserID, PayloadData.Content)
		// TODO: Upon incoming message, the hub finds the client, and uses the write pump function to write the message
		// Go over the active clients
		// Find the corresponding receiver ID in the active connections
		// Write message to their send channel
		// Write this information to their socket connection channel
	}
}

func (Client *Client) WritePump(){
	// TODO: Implement the write pump
	// Its purpose is to utilize the socket channel to write data to from the sendchannel of the target client

	// 1. Constantly read from send channel
	// 2. Write from send channel to socket connection, using the WriteMessage function\
	for{
		msg := <-Client.SendChannel
		err := Client.Connection.WriteMessage(websocket.TextMessage,msg)
		if err != nil{
			log.Println("An error occurred")
			log.Println(err)
		}
	}
}

func (Client *Client) ReadPump(){
	// TODO: Implement the Read Pump
	// Its purpose is to take notice of any user changes, i.e., disconnected, is typing, etc.
}

func SendMessageToClient(Hub *Hub, ReceiverID bson.ObjectID, Message string){
	for client := range Hub.ActiveClients{
		if client.UserID == ReceiverID{
			client.SendChannel <- []byte(Message)
		}
	}
}
