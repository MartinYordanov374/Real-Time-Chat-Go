package HelperFunctions

import (
	"RealTimeChatApp/Backend/Redis"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserSession(GinContext *gin.Context) (Redis.Session, error) {
	SessionCookie, err := GinContext.Cookie("SessionID")
	RedisSession, err := Redis.Client.Get(context.Background(), SessionCookie).Result()
	if err != nil {
		log.Println(err)
		log.Println("The session is inactive")
		return "", err
	}
	var SessionData Redis.Session
	RedisData := []byte(RedisSession)
	err = json.Unmarshal(RedisData, &SessionData)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return SessionData, nil
}