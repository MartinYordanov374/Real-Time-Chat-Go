package HandlerFunctions

import (
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/GlobalVariables"
	"context"
	"log"
	"RealTimeChatApp/Backend/Mongo"

)
func Login(context *gin.Context){
	context.JSON(200, gin.H{
		"message": "This is the logind endpoint placeholder",
	})
}

func Register(GinContext *gin.Context){
	newUser := MongoConfig.User{Username: "UserTemplate", HashedPassword: "nothashed", Email: "notmail"}
	var UserData MongoConfig.User
	GinContext.BindJSON(&UserData)
	GinContext.JSON(200, UserData)

	log.Println(UserData)
	res, error := GlobalVariables.MongoCollection.InsertOne(context.TODO(), newUser)

	if error != nil{
		log.Println(error)
	}else{
		log.Println(res.InsertedID)
	}

	GinContext.JSON(200, gin.H{
		"message": "This is the registration endpoint placeholder",
	})
}
