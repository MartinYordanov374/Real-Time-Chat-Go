package HelperFunctions

import (
	"RealTimeChatApp/Backend/Redis"
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserSession(GinContext *gin.Context) (Redis.Session, error) {
	SessionCookie, err := GinContext.Cookie("SessionID")
	RedisSession, err := Redis.Client.Get(context.Background(), SessionCookie).Result()
	if err != nil {
		log.Println(err)
		log.Println("The session is inactive")
		return Redis.Session{}, err
	}
	var SessionData Redis.Session
	RedisData := []byte(RedisSession)
	err = json.Unmarshal(RedisData, &SessionData)

	if err != nil {
		log.Println(err)
		return Redis.Session{}, err
	}

	return SessionData, nil
}
