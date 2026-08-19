package HandlerFunctions

import (
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/GlobalVariables"
	"context"
	"log"
	"RealTimeChatApp/Backend/Mongo"
	"regexp"
	"net/mail"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)
func Login(context *gin.Context){
	context.JSON(200, gin.H{
		"message": "This is the logind endpoint placeholder",
	})
}

func Register(GinContext *gin.Context){
	var UserData MongoConfig.User
	GinContext.BindJSON(&UserData)
	GinContext.JSON(200, UserData)

	if UsernameExists(UserData.Username) || EmailExists(UserData.Email){
		log.Println("This username or email is already taken!")
	}else{
		if ValidateUsername(UserData.Username){
			if ValidateEmail(UserData.Email){
				if ValidatePassword(UserData.Password){
					// TODO: Hash the password
					HashedPassword := HashPassword(UserData.Password)
					newUser := MongoConfig.User{Username: UserData.Username, Password: HashedPassword, Email: UserData.Email}
					_, error := GlobalVariables.MongoUsersCollection.InsertOne(context.TODO(), newUser)
					if error != nil{
						log.Println(error)
					}else{
						log.Println("User registered Successfully")
					}
				}else{
					log.Println("Invalid password")
				}
			}else{
				log.Println("Invalid email")
			}
		}else{
				log.Println("Invalid Username")
		}
	}
}


func ValidateUsername(Username string) bool{
	if len(Username) >= 2{
		UsernameRegex, _ := regexp.Compile("^[a-zA-Z]{2,}$")

		ValidUsername := UsernameRegex.MatchString(Username)
		// TODO: Check if the username is already taken

		if ValidUsername {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}

func ValidatePassword(Password string) bool{
	// TODO: Move all regexes to a seperate file
	if len(Password) >= 15{
		AtLeastOneLowerCaseRegex:= regexp.MustCompile(`[a-z]`)
		AtLeastOneUpperCaseRegex:= regexp.MustCompile(`[A-Z]`)
		AtLeastOneDigitRegex := regexp.MustCompile(`[\\d]`)
		SpecialRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

		PasswordContainsLowerCase := AtLeastOneLowerCaseRegex.MatchString(Password)
		PasswordContainsUpperCase := AtLeastOneUpperCaseRegex.MatchString(Password)
		PasswordContainsDigit := AtLeastOneDigitRegex.MatchString(Password)
		PasswordHasSpecialCharacter := SpecialRegex.MatchString(Password)


		if (PasswordContainsLowerCase && PasswordContainsUpperCase && PasswordContainsDigit && PasswordHasSpecialCharacter){
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}

func ValidateEmail(Email string) bool{
	// TODO: Use an already existing and a battle-proven regex for email verification
	//
	_, err := mail.ParseAddress(Email)
	if err != nil{
		return false
	}else{
		return true
	}

}
func HashPassword(Password string) string{
	// TODO: The password shoihuld be using a cost value from an env file.
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil{
		panic(err)
	}

	return string(HashedPassword)
}

func UsernameExists(Username string) bool{
	var user MongoConfig.User
	filter := bson.M{"username": Username}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil{
		log.Println(err)
		return false
	}else{
		return true
	}
}

func EmailExists(Email string) bool{
	var user MongoConfig.User
	filter := bson.M{"email": Email}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil{
		log.Println(err)
		return false
	}else{
		return true
	}

	return false
}
