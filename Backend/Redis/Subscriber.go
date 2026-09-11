package Redis

import (
	MongoConfig "RealTimeChatApp/Backend/Mongo"
	"RealTimeChatApp/Backend/WebSockets"
	"context"
	"encoding/json"
	"log"
)

func RedisMessageSubscribeHandler(Hub *WebSockets.Hub) {
	sub := Client.Subscribe(context.Background(), "Message")
	defer sub.Close()
	for {
		msg, err := sub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		var PayloadData MongoConfig.Message
		err = json.Unmarshal([]byte(msg.Payload), &PayloadData)
		if err != nil {
			log.Println(err)
		}

		WebSockets.SendMessageToClient(Hub, PayloadData)
	}
}
