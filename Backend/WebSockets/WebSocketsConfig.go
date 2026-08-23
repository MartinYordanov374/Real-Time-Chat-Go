package WebSockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"RealTimeChatApp/Backend/Redis"
	"context"
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

	SessionCookie, err := GinContext.Cookie("SessionID")
	if err != nil{
		log.Println(err)
		return
	}

	// TODO: This is repeating code, move it to a seperate file
	RedisSession, err := Redis.Client.Get(context.TODO(), SessionCookie).Result()
	if err != nil{
		log.Println(err)
		return
	}else{
		var SessionData Redis.Session;
		RedisData := []byte(RedisSession)
		err := json.Unmarshal(RedisData, &SessionData)

		if err != nil {
			log.Println(err)
			return
		}

		UserID := SessionData.UserID
		Client := &Client{
			UserID: UserID,
			Connection: SocketConnection,
			SendChannel: make(chan []byte),
		}

		log.Println(Client)
		// TODO: Find a way to reference the target hub here without import cycles
		Handler.Hub.RegisterClient(Client)
		log.Println(Handler.Hub)
	}
}
