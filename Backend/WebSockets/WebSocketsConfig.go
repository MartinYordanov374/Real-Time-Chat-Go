package WebSockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

var ConnectionUpgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(*http.Request) bool{
		return true
	},
}


func HandleWebSocketConnection(GinContext *gin.Context){
	_, err := ConnectionUpgrader.Upgrade(GinContext.Writer, GinContext.Request, nil)

	if err != nil {
		log.Println("Connection upgrade error: ", err)
		return
	}

	log.Println("Upgraded to websocket!")
}
