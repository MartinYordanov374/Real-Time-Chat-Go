package Middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"RealTimeChatApp/Backend/Redis"
	"context"
)
func AuthMiddleware(GinContext *gin.Context) bool{
	SessionCookie, err := GinContext.Cookie("SessionID")
	if err != nil{
		GinContext.String(http.StatusNotFound, "Cookie missing")
		return false
	}
	_, RedisResultError := Redis.Client.Get(context.TODO(), SessionCookie).Result()
	if RedisResultError != nil{
		GinContext.String(http.StatusNotFound, "This session does not exist in redis")
		return false

	}else{
		GinContext.String(http.StatusOK, "This session exists in redis")
		return true
	}
}
