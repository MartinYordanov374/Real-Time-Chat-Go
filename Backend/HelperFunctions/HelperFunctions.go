package HelperFunctions

import(
	"regexp"
	"strings"
	"net/mail"
	"golang.org/x/crypto/bcrypt"
	"RealTimeChatApp/Backend/GlobalVariables"
	"RealTimeChatApp/Backend/Mongo"
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

func SetSessionCookie(GinContext *gin.Context, SessionID string){
	// TODO: Make the cookie last as long as the session, i.e., create a global variable for this
	GinContext.SetCookie("SessionID", SessionID, GlobalVariables.CookieExpirationSeconds, "/", "localhost", false, false)
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

func UserExistsByID(UserID bson.ObjectID) bool{
	var user MongoConfig.User
	filter := bson.M{"_id": UserID}
	err := GlobalVariables.MongoUsersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false
	}else{
		return true
	}
}

// TODO: Move Chat functions to a seperate file
func ChatExistsBetweenUsers(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool{
	var chat MongoConfig.Chat
	filter := bson.M{"creator_id": SenderID, "receiver_id": ReceiverID}
	err := GlobalVariables.MongoChatsCollection.FindOne(context.TODO(), filter).Decode(&chat)

	if err != nil {
		log.Println("Chat between those users does not exist, creating chat...")
		return false
	}else{
		log.Println("The chat between those users exists, sending message...")
		return true
	}

}

func CreateChatObject(CreatorID bson.ObjectID, ReceiverID bson.ObjectID){
	newChat := MongoConfig.Chat{
		Messages: []bson.ObjectID{},
		CreatorID: CreatorID,
		ReceiverID: ReceiverID,
		CreationDate: time.Now()}

	_, err := GlobalVariables.MongoChatsCollection.InsertOne(context.TODO(), newChat)

	if err != nil {
		log.Println(err)
	}else{
		log.Println("Successfully created chat between the users")
	}
}

func CreateMessageObject(SenderID bson.ObjectID, ChatID bson.ObjectID, Content string){
	newMessage := MongoConfig.Message{
		ChatID: ChatID,
		TextContent: Content,
		TimeStamp: time.Now(),
		SenderID: SenderID,
	}

	_, err := GlobalVariables.MongoMessagesCollection.InsertOne(context.TODO(), newMessage)
	if err != nil{
		log.Println(err)
	}else{
		log.Println("Message object created")
	}
}

func RetrieveChatID(CreatorID bson.ObjectID, ReceiverID bson.ObjectID) bson.ObjectID{
	// TODO: Make those filters bi-directional
	var TargetChat MongoConfig.Chat;
	filter := bson.M{"creator_id": CreatorID, "receiver_id": ReceiverID}
	err := GlobalVariables.MongoChatsCollection.FindOne(context.TODO(), filter).Decode(&TargetChat)
	if err != nil{
		log.Println(err)
		return bson.NilObjectID
	}else{
		return TargetChat.ID
	}
}

func ChatRequestSent(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool{
	return false
}

func SendChatRequest(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool{
	return false
}

func IsChatRequestAccepted(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool{
	return false
}
