package HelperFunctions

import (
	"RealTimeChatApp/Backend/GlobalVariables"
	"RealTimeChatApp/Backend/Mongo"
	"context"
	"log"
	"net/mail"
	"regexp"
	"strings"
	"time"
	"RealTimeChatApp/Backend/Redis"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
)

func SetSessionCookie(GinContext *gin.Context, SessionID string) {
	GinContext.SetCookie("SessionID", SessionID, GlobalVariables.CookieExpirationSeconds, "/", "localhost", false, false)
}

func ValidateUsername(Username string) bool {
	TrimmedUsername := strings.TrimSpace(Username)
	if len(TrimmedUsername) >= 2 {
		UsernameRegex, _ := regexp.Compile("^[a-zA-Z]{2,}$")

		ValidUsername := UsernameRegex.MatchString(TrimmedUsername)
		if ValidUsername {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ValidatePassword(Password string) bool {
	// TODO: Move all regexes to a seperate file
	TrimmedPassword := strings.TrimSpace(Password)
	if len(TrimmedPassword) >= 15 {
		AtLeastOneLowerCaseRegex := regexp.MustCompile(`[a-z]`)
		AtLeastOneUpperCaseRegex := regexp.MustCompile(`[A-Z]`)
		AtLeastOneDigitRegex := regexp.MustCompile(`[\\d]`)
		SpecialRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)

		PasswordContainsLowerCase := AtLeastOneLowerCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsUpperCase := AtLeastOneUpperCaseRegex.MatchString(TrimmedPassword)
		PasswordContainsDigit := AtLeastOneDigitRegex.MatchString(TrimmedPassword)
		PasswordHasSpecialCharacter := SpecialRegex.MatchString(TrimmedPassword)

		if PasswordContainsLowerCase && PasswordContainsUpperCase && PasswordContainsDigit && PasswordHasSpecialCharacter {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ValidateEmail(Email string) bool {
	TrimmedEmai := strings.TrimSpace(Email)
	_, err := mail.ParseAddress(TrimmedEmai)
	if err != nil {
		return false
	} else {
		return true
	}
}

func HashPassword(Password string) string {
	// TODO: The password shoihuld be using a cost value from an env file.
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(HashedPassword)
}

func UsernameExists(Username string) bool {
	// TODO: Rename to UserExists
	var user MongoConfig.User
	TrimmedUsername := strings.TrimSpace(Username)
	filter := bson.M{"username": TrimmedUsername}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func EmailExists(Email string) bool {
	var user MongoConfig.User
	TrimmedEmail := strings.TrimSpace(Email)
	filter := bson.M{"email": TrimmedEmail}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}

	return false
}

func UserExistsByID(UserID bson.ObjectID) bool {
	var user MongoConfig.User
	filter := bson.M{"_id": UserID}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false
	} else {
		return true
	}
}

// TODO: Move Chat functions to a seperate file
func ChatExistsBetweenUsers(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool {
	var chat MongoConfig.Chat
	filter := bson.M{"creator_id": SenderID, "receiver_id": ReceiverID}
	err := GlobalVariables.MongoChatsCollection.FindOne(context.TODO(), filter).Decode(&chat)

	if err != nil {
		return false
	} else {
		return true
	}
}

func CreateChatObject(CreatorID bson.ObjectID, ReceiverID bson.ObjectID) {
	newChat := MongoConfig.Chat{
		CreatorID:    CreatorID,
		ReceiverID:   ReceiverID,
		CreationDate: time.Now(),
	}

	_, err := GlobalVariables.MongoChatsCollection.InsertOne(context.TODO(), newChat)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully created chat between the users")
	}
}

func CreateMessageObject(SenderID bson.ObjectID, ChatID bson.ObjectID, Content string) MongoConfig.Message {
	newMessage := MongoConfig.Message{
		ID:          bson.NewObjectID(),
		ChatID:      ChatID,
		TextContent: Content,
		TimeStamp:   time.Now(),
		SenderID:    SenderID,
	}

	_, err := GlobalVariables.MongoMessagesCollection.InsertOne(context.TODO(), newMessage)
	if err != nil {
		log.Println(err)
	}
	return newMessage
}

func RetrieveChatID(CreatorID bson.ObjectID, ReceiverID bson.ObjectID) bson.ObjectID {
	// TODO: Make those filters bi-directional
	var TargetChat MongoConfig.Chat
	filter := bson.M{"creator_id": CreatorID, "receiver_id": ReceiverID}
	err := GlobalVariables.MongoChatsCollection.FindOne(context.TODO(), filter).Decode(&TargetChat)
	if err != nil {
		log.Println(err)
		return bson.NilObjectID
	} else {
		return TargetChat.ID
	}
}

func ChatRequestSent(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool {
	// TODO: Check if the request object exists
	filter := bson.M{"sender_id": SenderID, "receiver_id": ReceiverID}
	var TargetRequest MongoConfig.Request
	err := GlobalVariables.MongoRequestsCollection.FindOne(context.TODO(), filter).Decode(&TargetRequest)
	if err != nil {
		log.Println(err)
		return false
	} else {
		log.Println(TargetRequest)
		return true
	}
}

func SendChatRequest(SenderID bson.ObjectID, ReceiverID bson.ObjectID) {
	// TODO: Create Request Object
	if !ChatRequestSent(SenderID, ReceiverID) {
		NewRequest := MongoConfig.Request{
			SenderID:   SenderID,
			ReceiverID: ReceiverID,
			Status:     MongoConfig.RequestPending,
			TimeStamp:  time.Now(),
		}
		_, err := GlobalVariables.MongoRequestsCollection.InsertOne(context.TODO(), NewRequest)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Connection Request sent!")
		}
	} else {
		log.Println("You have already sent this user a message request!")
	}
}

func GetChatRequestStatus(SenderID bson.ObjectID, ReceiverID bson.ObjectID) MongoConfig.RequestStatus {

	filter := bson.M{"sender_id": SenderID, "receiver_id": ReceiverID}
	var TargetRequest MongoConfig.Request

	err := GlobalVariables.MongoRequestsCollection.FindOne(context.TODO(), filter).Decode(&TargetRequest)
	if err != nil {
		log.Println(err)
		return MongoConfig.RequestError
	}
	return TargetRequest.Status
}

func AcceptChatRequest(RequestID bson.ObjectID) {
	// TODO: This function shall handle a user's response to a request
	// 1. If the request is rejected, delete all conversation and correspondingb messages with the sender
	// 2. If the request is approved, the chat remains and the sender can send more messages than just one.

	// TODO: Verify that the requesting user ID is the same as RECEIVER ID for the said request
	filter := bson.M{"_id": RequestID}
	Update := bson.M{
		"$set": bson.M{
			"request_status": MongoConfig.RequestAccepted,
		},
	}

	_, err := GlobalVariables.MongoRequestsCollection.UpdateOne(context.TODO(), filter, Update)
	if err != nil {
		log.Println(err)
	}
}

func RejectChatRequest(RequestID bson.ObjectID) {
	// TODO: Verify that the requesting user ID is the same as RECEIVER ID for the said request
	filter := bson.M{"_id": RequestID}
	Update := bson.M{
		"$set": bson.M{
			"request_status": MongoConfig.RequestRejected,
		},
	}

	_, err := GlobalVariables.MongoRequestsCollection.UpdateOne(context.TODO(), filter, Update)
	if err != nil {
		log.Println(err)
	}
	// TODO: Delete the respective conversation and all associated messages upon rejection
	DeleteRejectedRequestChat(RequestID)
}

func DeleteRejectedRequestChat(RequestID bson.ObjectID) {
	var request MongoConfig.Request
	filter := bson.M{"_id": RequestID}
	err := GlobalVariables.MongoRequestsCollection.FindOne(context.TODO(), filter).Decode(&request)
	if err != nil {
		log.Println(err)
	}

	var chat MongoConfig.Chat
	chatFilter := bson.M{
		"$or": []bson.M{
			{
				"creator_id":  request.SenderID,
				"receiver_id": request.ReceiverID,
			},
			{
				"creator_id":  request.ReceiverID,
				"receiver_id": request.SenderID,
			},
		},
	}

	ChatErr := GlobalVariables.MongoChatsCollection.FindOne(context.TODO(), chatFilter).Decode(&chat)

	if ChatErr != nil {
		log.Println(ChatErr)
	}

	messageFilter := bson.M{
		"chat_id": chat.ID,
	}

	_, MessagesErr := GlobalVariables.MongoMessagesCollection.DeleteMany(context.TODO(), messageFilter)

	if MessagesErr != nil {
		log.Println(MessagesErr)
	}

	_, ChatDelErr := GlobalVariables.MongoChatsCollection.DeleteOne(context.TODO(), bson.M{"_id": chat.ID})
	if ChatDelErr != nil {
		log.Println(ChatDelErr)
	}

	_, RequestDelErr := GlobalVariables.MongoRequestsCollection.DeleteOne(context.TODO(), bson.M{"_id": RequestID})

	if RequestDelErr != nil {
		log.Println(RequestDelErr)
	}

}

func CacheMessageRedis(Message MongoConfig.Message){
	// TODO: If a message gets edited or deleted, the cache for the said conversation should be destroyed entireloy
	key := "chat:"+Message.ChatID.Hex()+":messages"

	MarshaledMessage, err := json.Marshal(Message)
	if err != nil {
		log.Println(err)
		return
	}
	// TODO: Expire the stored messages after, say, 12 h
	pipeline := Redis.Client.Pipeline()
	pipeline.LPush(context.TODO(), key, MarshaledMessage)
	pipeline.LTrim(context.TODO(), key, 0, 99)

	_, err = pipeline.Exec(context.TODO())
	if err != nil {
		log.Println(err)
	}
}

func GetCachedMessages(ChatID bson.ObjectID) ([]MongoConfig.Message, error){
	// TODO: If the messages do not exist in cache,
	// then fetch them from DB and save them via CacheMessageRedis
	key := "chat:"+ChatID.Hex()+":messages"
	msgs, err := Redis.Client.LRange(context.TODO(), key, 0, 99).Result()

	if err != nil{
		log.Println(err)
		// NOTE: This should work because if a conversation between users exist, then it should have at least one message given that nobody deleted the original message
		return nil, err
	}

	if len(msgs) == 0{
		// TODO: Fetch the DB later
		return []MongoConfig.Message{}, nil
	}

	TargetMessages := make([]MongoConfig.Message, 0, len(msgs))
	for _, Message := range msgs {
		var TargetMessage MongoConfig.Message
		err := json.Unmarshal([]byte(Message), &TargetMessage)
		if err != nil {
			continue
		}
		TargetMessages = append(TargetMessages, TargetMessage)
	}
	return TargetMessages, nil
}


func RetrieveAllUserChats(UserID bson.ObjectID) []MongoConfig.Chat{
	// TODO: Get rid of the lookup. Fetch the conversations. Fetch the last 50 messages only if they are not already within Redis cache.
	// If they are not cached in Redis, cache them. Only load further messages if you need to find them. 50 per request. No lookup.

	// 1. Find all chats where UserID is either sender or creator
	var TargetChats []MongoConfig.Chat
	UserChatsFilter := bson.M{
		"$or": []bson.M{
			{"creator_id": UserID},
			{"receiver_id": UserID},
		}}

	cursor, err := GlobalVariables.MongoChatsCollection.Find(context.TODO(), UserChatsFilter)

	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &TargetChats)
	if err != nil{
		log.Println(err)
	}
	// Retrieving Messages
	for idx, TargetChat := range TargetChats{
		res, err := GetCachedMessages(TargetChat.ID)
		if err != nil {
			log.Println(err)
		}
		log.Println(res)
		TargetChats[idx].Messages = res
	}

	return TargetChats
	// 2. Check if messages with the corersponding Chat IDs are already cached in Redis
	// 3. If not, fetch the mongoDB for the last 100 messages and stores them in Redis.
	// 4. If the users scrolls up to the 100th message, fetch the next 100 messages from and store them in Redis.
	// 5. Redis shall contain no more than 100 of the latest messages.
}
