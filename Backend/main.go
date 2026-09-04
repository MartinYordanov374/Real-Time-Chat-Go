package main

import(
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/Handlers"
	"RealTimeChatApp/Backend/Mongo"
	"RealTimeChatApp/Backend/Middlewares"
	"RealTimeChatApp/Backend/WebSockets"
)

func main(){

	MongoConfig.ConnectToMongo()

	router := gin.Default()

	// TODO: Make the login and register endpoints unavailable for logged registers
	// TODO: Create a logout endpoint

	WebSocketsHub := WebSockets.NewHub()
	WebSocketsHandler := WebSockets.CreateHandler(WebSocketsHub)
	router.POST("/login", Middlewares.CORSMiddleware(), HandlerFunctions.Login)

	router.POST("/register", Middlewares.CORSMiddleware(), HandlerFunctions.Register)

	// TODO; Implement logout
	router.POST("/logout", Middlewares.CORSMiddleware(), Middlewares.AuthMiddleware(), HandlerFunctions.Logout)

	router.GET("/RetrieveChat/:ChatID", Middlewares.CORSMiddleware(), Middlewares.AuthMiddleware(), HandlerFunctions.RetrieveChat)
	router.POST("/SendMessage/:ReceiverID", Middlewares.CORSMiddleware(), Middlewares.AuthMiddleware(), HandlerFunctions.SendMessage)
	router.POST("/InviteUserToGroupChat/:UserID/:ChatID", Middlewares.CORSMiddleware(), Middlewares.AuthMiddleware(), HandlerFunctions.InviteUserToGroupChat)
	router.POST("/AcceptChatRequest/:RequestID", Middlewares.CORSMiddleware(), Middlewares.AuthMiddleware(), HandlerFunctions.AcceptChatRequest)
	router.POST("/RejectChatRequest/:RequestID", Middlewares.CORSMiddleware(), Middlewares.AuthMiddleware(), HandlerFunctions.RejectChatRequest)
	router.GET("/ws", Middlewares.CORSMiddleware(), Middlewares.AuthMiddleware(), WebSocketsHandler.HandleWebSocketConnection)
	router.Run()
}
