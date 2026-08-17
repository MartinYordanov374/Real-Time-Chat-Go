package main

import(
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/Handlers"
)

func main(){
	router := gin.Default()
	router.POST("/login", HandlerFunctions.Login)

	router.POST("/register", HandlerFunctions.Register)

	router.Run()
}
