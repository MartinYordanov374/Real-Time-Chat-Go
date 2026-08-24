package WebSockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	go WritePump()
	go ReadPump(Handler.Hub)
}

func WritePump(){

}

func ReadPump(Hub *Hub){

}
