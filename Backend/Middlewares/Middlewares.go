package Middlewares

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/Redis"
	"encoding/json"
	"log"
)
func AuthMiddleware() gin.HandlerFunc{
	return func(GinContext *gin.Context){
		SessionCookie, err := GinContext.Cookie("SessionID")
		if err != nil{
			GinContext.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"Cookie Missing"})
			return
		}
		RedisSession, RedisResultError := Redis.Client.Get(context.TODO(), SessionCookie).Result()
		if RedisResultError != nil{
			GinContext.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"This session does not exist in redis"})
			return

		}else{
			var SessionData Redis.Session;
			RedisData := []byte(RedisSession)
			err := json.Unmarshal(RedisData, &SessionData)

			if err != nil {
				log.Println(err)
				return
			}
			GinContext.Set("UserID", SessionData.UserID)
			GinContext.Next()
		}
	}
}
