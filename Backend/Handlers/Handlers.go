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
	"RealTimeChatApp/Backend/Redis"
	"github.com/google/uuid"
	"time"
	"encoding/json"
	"net/http"
)

// TODO: Add more descriptive error messages, i.e., tell what the requirements for a pass and username are.
func Login(GinContext *gin.Context){

	var UserData MongoConfig.User
	GinContext.BindJSON(&UserData)

	if ValidateUsername(UserData.Username){
		if UsernameExists(UserData.Username){
			var TargetUser MongoConfig.User
			filter := bson.M{"username":UserData.Username}
			err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&TargetUser)
			if err != nil{
				GinContext.JSON(500, gin.H{"message": "Internal Server Error"})
				return
			}

			hashErr := bcrypt.CompareHashAndPassword([]byte(TargetUser.Password), []byte(UserData.Password))
			if hashErr != nil{
				GinContext.JSON(401, gin.H{"message": "Wrong Password"})
			}else{
				SessionID := uuid.New().String()
				SessionData, SessionError := json.Marshal(Redis.Session{SessionID: SessionID, UserID:TargetUser.ID})
				if SessionError != nil{
					log.Println(SessionError)
				}
				err := Redis.Client.Set(context.TODO(), SessionID, SessionData, 1*time.Minute).Err()
				if err != nil{
					log.Println(err)
				}

				SetSessionCookie(GinContext, SessionID)
				GinContext.JSON(200, gin.H{"message": "Login Successful"})

			}
		}else{
				GinContext.JSON(404, gin.H{"message": "Such a user does not exist!"})
		}
	}else{
		GinContext.JSON(401, gin.H{"message": "This username does not meet the username requirements."})

	}
}

func Register(GinContext *gin.Context){
	// TODO: If the username is unique, then conbsider removing the email field
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

func TestSession(GinContext *gin.Context){
	// TODO:
	// 1. Rename this to AuthMiddleware
	// 2. Use this function as a middleware for the planned endpoints
	SessionCookie, err := GinContext.Cookie("SessionID")
	if err != nil{
		GinContext.String(http.StatusNotFound, "Cookie missing")
		return
	}

	log.Println(SessionCookie)
	_, RedisResultError := Redis.Client.Get(context.TODO(), SessionCookie).Result()
	if RedisResultError != nil{
		GinContext.String(http.StatusNotFound, "This session does not exist in redis")

	}else{
		GinContext.String(http.StatusOK, "This session exists in redis")
	}
}

func SetSessionCookie(GinContext *gin.Context, SessionID string){
	// TODO: Make the cookie last as long as the session, i.e., create a global variable for this
	GinContext.SetCookie("SessionID", SessionID, 60, "/", "localhost", false, false)
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
	// TODO: Rename to UserExists
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
