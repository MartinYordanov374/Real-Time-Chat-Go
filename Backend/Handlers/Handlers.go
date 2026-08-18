package HandlerFunctions

import (
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/GlobalVariables"
	"context"
	"log"
	"RealTimeChatApp/Backend/Mongo"
	"go.mongodb.org/mongo-driver/v2/bson"

)
func Login(context *gin.Context){
	context.JSON(200, gin.H{
		"message": "This is the logind endpoint placeholder",
	})
}

func Register(GinContext *gin.Context){
	// TODO: Add data validations and sanitization
	// TODO: Hash the password
	var UserData MongoConfig.User
	GinContext.BindJSON(&UserData)
	GinContext.JSON(200, UserData)

	newUser := MongoConfig.User{Username: UserData.Username, HashedPassword: UserData.HashedPassword, Email: UserData.Email}

	res, error := GlobalVariables.MongoUsersCollection.InsertOne(context.TODO(), newUser)

	if error != nil{
		log.Println(error)
	}else{
		log.Println(res.InsertedID)
	}
}
