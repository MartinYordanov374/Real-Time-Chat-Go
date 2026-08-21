package HandlerFunctions

import (
	"github.com/gin-gonic/gin"
	"RealTimeChatApp/Backend/GlobalVariables"
	"context"
	"log"
	"RealTimeChatApp/Backend/Mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
 	"RealTimeChatApp/Backend/Redis"
	"RealTimeChatApp/Backend/HelperFunctions"
	"github.com/google/uuid"
	"encoding/json"
)

// TODO: Add more descriptive error messages, i.e., tell what the requirements for a pass and username are.
func Login(GinContext *gin.Context){

	var UserData MongoConfig.User
	GinContext.BindJSON(&UserData)

	if HelperFunctions.ValidateUsername(UserData.Username){
		if HelperFunctions.UsernameExists(UserData.Username){
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
				err := Redis.Client.Set(context.TODO(), SessionID, SessionData, GlobalVariables.SessionDuration).Err()
				if err != nil{
					log.Println(err)
				}

				HelperFunctions.SetSessionCookie(GinContext, SessionID)
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

	if HelperFunctions.UsernameExists(UserData.Username) || HelperFunctions.EmailExists(UserData.Email){
			GinContext.JSON(200, gin.H{
				"message": "This username or email is already taken!"})
	}else{
		if HelperFunctions.ValidateUsername(UserData.Username){
			if HelperFunctions.ValidateEmail(UserData.Email){
				if HelperFunctions.ValidatePassword(UserData.Password){
					HashedPassword := HelperFunctions.HashPassword(UserData.Password)
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


func Test(GinContext *gin.Context){
	GinContext.JSON(200, gin.H{"message": "Success"})
}
