package HelperFunctions

import (
	"RealTimeChatApp/Backend/Redis"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserSession(GinContext *gin.Context) (string, error) {
	SessionCookie, err := GinContext.Cookie("SessionID")
	RedisSession, err := Redis.Client.Get(context.Background(), SessionCookie).Result()
	if err != nil {
		log.Println(err)
		log.Println("The session is inactive")
		return "", err
	}
	return RedisSession, nil
}
