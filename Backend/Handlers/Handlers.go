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
				SessionData, SessionError := json.Marshal(Redis.Session{UserID:TargetUser.ID})
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

// TODO: The below functions are accessible to logged in users only
// TODO: Take care of authorization also
// TODO: Each sent message that is not read by the receiver shall land in the inbox
// TODO: The inbox shall also include history and the received chat requests
func SendMessage(GinContext *gin.Context){
	// TODO: Follow the steps below
	// 1. Get the sender's data via the session cookie
	SessionCookie, err := GinContext.Cookie("SessionID")

	if err != nil{
		log.Println(err)
	}else{
		log.Println(SessionCookie)
		RedisSession, err := Redis.Client.Get(context.TODO(), SessionCookie).Result()
		if err != nil{
			log.Println(err)
		}else{
			var SessionData Redis.Session;
			RedisData := []byte(RedisSession)
			err := json.Unmarshal(RedisData, &SessionData)

			if err != nil {
				log.Println(err)
				return
			}

			SenderID := SessionData.UserID
			log.Println("Sender: ", SenderID)
			ReceiverID := GinContext.Param("ReceiverID")
			log.Println("Receiver: ", ReceiverID)
			ConvertedReceiverID, err := bson.ObjectIDFromHex(ReceiverID)
			if err != nil {
				log.Println(err)
				return
			}else{
				if HelperFunctions.UserExistsByID(ConvertedReceiverID){
					log.Println("The receiver user exists")
					if HelperFunctions.ChatExistsBetweenUsers(SenderID, ConvertedReceiverID){
						// TODO: Create message object with the chat ID
					}else{
						// TODO: Create chat between the users and a sender message ID to this chat
						// TODO: Create a request object from sender to reciever
						// NOTE: No more messages can be sent until the receiver accepts the request
						HelperFunctions.CreateChatObject(SenderID, ConvertedReceiverID)

					}
				}else{
					log.Println("The receiver user does not exist")
					return
				}
			}
		}
	}
	// 2. Validate whether the receiver user exists
	// 3. Validate whether a chat between the sender and the receiver exists, if not send a request to the receiver
	// 3.1. Send the sender message regardless of whether the receiver accepts the chat invitation or not
	// 3.2. If the receiver rejects the invitation, delete the conversation along with the messages.

	// TODO: Create a CreateConversation helper function
}


func RespondToChatRequest(GinContext *gin.Context){
	// TODO: This function shall handle a user's response to a request
	// 1. If the request is rejected, delete all conversation and correspondingb messages with the sender
	// 2. If the request is approved, the chat remains and the sender can send more messages than just one.
}

func Logout(GinContext *gin.Context){
	// TODO: This function deleted the session cookie and removes the session from the Redis session storage
}


func InviteUserToGroupChat(GinContext *gin.Context){
	// TODO: This function shall invite a user to join an already existing chat between two or more users.
	// 1. Validate that the receiver user exists
	// 2. Send them a request and wait for their response
}

func RetrieveChat(GinContext *gin.Context){
	// TODO: This function shall fetch the chat between the requesting user and the specified user.
	// TODO: Retrieve the most recent, i.e., 50 or 100 messages from the chat from Redis if available
}
