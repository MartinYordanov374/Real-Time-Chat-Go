package main

import (
	HandlerFunctions "RealTimeChatApp/Backend/Handlers"
	"RealTimeChatApp/Backend/Middlewares"
	MongoConfig "RealTimeChatApp/Backend/Mongo"
	"RealTimeChatApp/Backend/Redis"
	"RealTimeChatApp/Backend/WebSockets"

	"github.com/gin-gonic/gin"
)

func main() {

	MongoConfig.ConnectToMongo()

	router := gin.Default()
	router.Use(Middlewares.CORSMiddleware())
	// TODO: Make the login and register endpoints unavailable for logged registers
	// TODO: Create a logout endpoint

	WebSocketsHub := WebSockets.NewHub()
	WebSocketsHandler := WebSockets.CreateHandler(WebSocketsHub)
	go Redis.RedisMessageSubscribeHandler(WebSocketsHub)
	router.POST("/login", HandlerFunctions.Login)
	router.POST("/register", HandlerFunctions.Register)
	router.POST("/logout", Middlewares.AuthMiddleware(), HandlerFunctions.Logout)
	router.GET("/RetrieveChat/:ChatID", Middlewares.AuthMiddleware(), HandlerFunctions.RetrieveChat)
	router.POST("/SendMessage/:ReceiverID", Middlewares.AuthMiddleware(), HandlerFunctions.SendMessage)
	router.POST("/InviteUserToGroupChat/:UserID/:ChatID", Middlewares.AuthMiddleware(), HandlerFunctions.InviteUserToGroupChat)
	router.POST("/AcceptChatRequest/:RequestID", Middlewares.AuthMiddleware(), HandlerFunctions.AcceptChatRequest)
	router.POST("/RejectChatRequest/:RequestID", Middlewares.AuthMiddleware(), HandlerFunctions.RejectChatRequest)
	router.GET("/ws", Middlewares.AuthMiddleware(), WebSocketsHandler.HandleWebSocketConnection)

	router.GET("/GetAllChats", Middlewares.AuthMiddleware(), HandlerFunctions.RetrieveAllUserChats)
	router.GET("/GetCurrentUserData", Middlewares.AuthMiddleware(), HandlerFunctions.GetCurrentUserData)

	router.Run()
}
