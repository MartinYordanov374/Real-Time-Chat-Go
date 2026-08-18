package HandlerFunctions

import (
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/GlobalVariables"
	"context"
	"log"
	"RealTimeChatApp/Backend/Mongo"
	"regexp"
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

	ValidateUsername(UserData.Username)

	res, error := GlobalVariables.MongoUsersCollection.InsertOne(context.TODO(), newUser)

	if error != nil{
		log.Println(error)
	}else{
		log.Println(res.InsertedID)
	}
}


func ValidateUsername(Username string) bool{
	if len(Username) >= 2{
		UsernameRegex, _ := regexp.Compile("^[a-zA-Z]{2,}$")

		ValidUsername := UsernameRegex.MatchString(Username)

		if ValidUsername {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
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

func UsernameExists(Username string) bool{
	return false
}

func EmailExists(Email string) bool{
	return false
}
