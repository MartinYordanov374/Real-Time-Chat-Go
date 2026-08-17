package main

import "github.com/gin-gonic/gin"

func main(){
	router := gin.Default()
	router.POST("/login", func(context *gin.Context){
		context.JSON(200, gin.H{
			"message": "This is the login endpoint placeholder",
		})
	})

	router.POST("/register", func(context *gin.Context){
		context.JSON(200, gin.H{
			"message": "This is the registration endpoint",
		})
	})
	router.Run()
}
