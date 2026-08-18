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
	// TODO: Add data validations and sanitization
	// TODO: Hash the password
	var UserData MongoConfig.User
	GinContext.BindJSON(&UserData)
	GinContext.JSON(200, UserData)

	newUser := MongoConfig.User{Username: UserData.Username, Password: UserData.Password, Email: UserData.Email}

	res, error := GlobalVariables.MongoUsersCollection.InsertOne(context.TODO(), newUser)

	if error != nil{
		log.Println(error)
	}else{
		log.Println(res.InsertedID)
	}
}


func ValidateUsername(Username string){
	// TODO: Validate username according to the below criteria:
	// 1. The username does not contain special characters
	// 2. The username is at least 2 characters long
	// 3. Trim string input to remove whitespaces
	// 4. Ensure that no other user is registered with the same username
}

func ValidatePassword(Password string){
	// TODO: Validate password according to the following criteria:
	// 1. The password is at least 15 characters long
	// 2. It includes at least one upper-case character
	// 3. It contains at least one special character
	// 4. Trim to remove whitespace
	// 5. At least one number
	// 6. No consecutive repeating characters
}

func ValidateEmail(Email string){
	// TODO: Use an already existing and a battle-proven regex for email verification
}

func HashPassword(Password string) string{
	// TODO: This function hashes the password only after the validations have passed successfully.
	HashedPassword := "Placeholder Value"
	return HashedPassword
}
