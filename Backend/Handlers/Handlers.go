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
	var RequestBody MongoConfig.Message;
	GinContext.BindJSON(&RequestBody)

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
			// TODO: Consider the opportunity for the user to send messages to themselves
			// Call this chat "Notes"
			// TODO: Implement Chat Titles
			SenderID := SessionData.UserID
			ReceiverID := GinContext.Param("ReceiverID")
			ConvertedReceiverID, err := bson.ObjectIDFromHex(ReceiverID)
			if err != nil {
				log.Println(err)
				return
			}else{
				if HelperFunctions.UserExistsByID(ConvertedReceiverID){
					if !HelperFunctions.ChatRequestSent(SenderID, ConvertedReceiverID) {
						HelperFunctions.SendChatRequest(SenderID, ConvertedReceiverID)
						HelperFunctions.CreateChatObject(SenderID, ConvertedReceiverID)
						CurrentChatID := HelperFunctions.RetrieveChatID(SenderID, ConvertedReceiverID)
						HelperFunctions.CreateMessageObject(SenderID, CurrentChatID, RequestBody.TextContent)
					}else{
						RequestStatus := HelperFunctions.GetChatRequestStatus(SenderID, ConvertedReceiverID)
						if RequestStatus == MongoConfig.RequestAccepted{
							CurrentChatID := HelperFunctions.RetrieveChatID(SenderID, ConvertedReceiverID)
							HelperFunctions.CreateMessageObject(SenderID, CurrentChatID, RequestBody.TextContent)
							// TODO: Create a struct for pub/sub message that includes the receiver ID and TextContent
							// TODO: Marshal that struct below and publish it
							RedisMessage := GlobalVariables.RedisMessage{SenderID, RequestBody.TextContent}
							MarshaledData, err := json.Marshal(RedisMessage)
							if err != nil {
								log.Println(err)
							}
							pubSubErr := Redis.Client.Publish(context.TODO(), "Message", MarshaledData).Err()
							if pubSubErr != nil{
								log.Println(pubSubErr)
							}
						}else if RequestStatus == MongoConfig.RequestPending{
							GinContext.JSON(202, gin.H{"message":"The request hasn't been answered yet. You can only send one message before a request is approved."})
						}else{
							GinContext.JSON(404, gin.H{"message":"In any other case the conversation is deleted"})
						}
					}
				}else{
					GinContext.JSON(404, gin.H{"message": "The receiver user does not exist"})
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


func AcceptChatRequest(GinContext *gin.Context){
	// TODO: This function shall handle a user's response to a request
	// 1. If the request is rejected, delete all conversation and correspondingb messages with the sender
	// 2. If the request is approved, the chat remains and the sender can send more messages than just one.
	RequestID := GinContext.Param("RequestID")
	ConvertedRequestID, err := bson.ObjectIDFromHex(RequestID)
	HelperFunctions.AcceptChatRequest(ConvertedRequestID)
	if err != nil {
		log.Println(err)
		return
	}
}

func RejectChatRequest(GinContext *gin.Context){
	RequestID := GinContext.Param("RequestID")
	ConvertedRequestID, err := bson.ObjectIDFromHex(RequestID)
	HelperFunctions.RejectChatRequest(ConvertedRequestID)
	if err != nil {
		log.Println(err)
		return
	}

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
