package main

import(
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/Handlers"
	"RealTimeChatApp/Backend/Mongo"
	"RealTimeChatApp/Backend/Middlewares"
)
func main(){

	MongoConfig.ConnectToMongo()

	router := gin.Default()

	router.POST("/login", HandlerFunctions.Login)

	router.POST("/register", HandlerFunctions.Register)

	router.GET("/test", Middlewares.AuthMiddleware(), HandlerFunctions.Test)

	router.Run()
}
