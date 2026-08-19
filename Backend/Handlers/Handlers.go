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
	"strings"
)
func Login(context *gin.Context){
	context.JSON(200, gin.H{
		"message": "This is the logind endpoint placeholder",
	})
}

func Register(GinContext *gin.Context){
	var UserData MongoConfig.User
	GinContext.BindJSON(&UserData)

	if UsernameExists(UserData.Username) || EmailExists(UserData.Email){
			GinContext.JSON(200, gin.H{
				"message": "This username or email is already taken!"})
	}else{
		if ValidateUsername(UserData.Username){
			if ValidateEmail(UserData.Email){
				if ValidatePassword(UserData.Password){
					// TODO: Hash the password
					HashedPassword := HashPassword(UserData.Password)
					newUser := MongoConfig.User{Username: UserData.Username, Password: HashedPassword, Email: UserData.Email}
					_, error := GlobalVariables.MongoUsersCollection.InsertOne(context.TODO(), newUser)
					if error != nil{
						GinContext.JSON(500, gin.H{
							"message": error})
					}else{
						GinContext.JSON(200, gin.H{
							"message": "User registered Successfully"})
					}
				}else{
					GinContext.JSON(401, gin.H{
						"message": "Invalid Password"})
					}
			}else{
				GinContext.JSON(401, gin.H{
					"message": "Invalid email"})
				}

		}else{
			GinContext.JSON(401, gin.H{
					"message": "Invalid Username"})
			}
		}
	}


func ValidateUsername(Username string) bool{
	TrimmedUsername := strings.TrimSpace(Username)
	if len(TrimmedUsername) >= 2{
		UsernameRegex, _ := regexp.Compile("^[a-zA-Z]{2,}$")

		ValidUsername := UsernameRegex.MatchString(TrimmedUsername)
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
	TrimmedPassword := strings.TrimSpace(Password)
	if len(TrimmedPassword) >= 15{
		AtLeastOneLowerCaseRegex:= regexp.MustCompile(`[a-z]`)
		AtLeastOneUpperCaseRegex:= regexp.MustCompile(`[A-Z]`)
		AtLeastOneDigitRegex := regexp.MustCompile(`[\\d]`)
		SpecialRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

		PasswordContainsLowerCase := AtLeastOneLowerCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsUpperCase := AtLeastOneUpperCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsDigit := AtLeastOneDigitRegex.MatchString(TrimmedPassword)
		PasswordHasSpecialCharacter := SpecialRegex.MatchString(TrimmedPassword)


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
	TrimmedEmai := strings.TrimSpace(Email)
	_, err := mail.ParseAddress(TrimmedEmai)
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
	TrimmedUsername := strings.TrimSpace(Username)
	filter := bson.M{"username": TrimmedUsername}
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
	TrimmedEmail := strings.TrimSpace(Email)
	filter := bson.M{"email": TrimmedEmail}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil{
		log.Println(err)
		return false
	}else{
		return true
	}

	return false
}
