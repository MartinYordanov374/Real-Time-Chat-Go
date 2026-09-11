package HelperFunctions

import (
	"RealTimeChatApp/Backend/GlobalVariables"
	MongoConfig "RealTimeChatApp/Backend/Mongo"
	"RealTimeChatApp/Backend/Redis"
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ChatExistsBetweenUsers(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool {
	var chat MongoConfig.Chat
	filter := bson.M{"creator_id": SenderID, "receiver_id": ReceiverID}
	err := GlobalVariables.MongoChatsCollection.FindOne(context.Background(), filter).Decode(&chat)

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

	_, err := GlobalVariables.MongoChatsCollection.InsertOne(context.Background(), newChat)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully created chat between the users")
	}
}

func CreateMessageObject(SenderID bson.ObjectID, ReceiverID bson.ObjectID, ChatID bson.ObjectID, Content string) MongoConfig.Message {
	newMessage := MongoConfig.Message{
		ID:          bson.NewObjectID(),
		ChatID:      ChatID,
		TextContent: Content,
		TimeStamp:   time.Now(),
		SenderID:    SenderID,
		ReceiverID:  ReceiverID,
	}

	_, err := GlobalVariables.MongoMessagesCollection.InsertOne(context.Background(), newMessage)
	if err != nil {
		log.Println(err)
	}
	return newMessage
}

func RetrieveChatID(CreatorID bson.ObjectID, ReceiverID bson.ObjectID) bson.ObjectID {
	var TargetChat MongoConfig.Chat
	filter := bson.M{
		"$or": []bson.M{
			{"creator_id": CreatorID, "receiver_id": ReceiverID},
			{"creator_id": ReceiverID, "receiver_id": CreatorID},
		},
	}
	err := GlobalVariables.MongoChatsCollection.FindOne(context.Background(), filter).Decode(&TargetChat)
	if err != nil {
		log.Println(err)
		return bson.NilObjectID
	} else {
		return TargetChat.ID
	}
}

func ChatRequestSent(SenderID bson.ObjectID, ReceiverID bson.ObjectID) bool {
	filter := bson.M{"sender_id": SenderID, "receiver_id": ReceiverID}
	var TargetRequest MongoConfig.Request
	err := GlobalVariables.MongoRequestsCollection.FindOne(context.Background(), filter).Decode(&TargetRequest)
	if err != nil {
		log.Println(err)
		return false
	} else {
		log.Println(TargetRequest)
		return true
	}
}

func SendChatRequest(SenderID bson.ObjectID, ReceiverID bson.ObjectID) {
	if !ChatRequestSent(SenderID, ReceiverID) {
		NewRequest := MongoConfig.Request{
			SenderID:   SenderID,
			ReceiverID: ReceiverID,
			Status:     MongoConfig.RequestPending,
			TimeStamp:  time.Now(),
		}
		_, err := GlobalVariables.MongoRequestsCollection.InsertOne(context.Background(), NewRequest)
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

	err := GlobalVariables.MongoRequestsCollection.FindOne(context.Background(), filter).Decode(&TargetRequest)
	if err != nil {
		log.Println(err)
		return MongoConfig.RequestError
	}
	return TargetRequest.Status
}

func AcceptChatRequest(RequestID bson.ObjectID) {
	filter := bson.M{"_id": RequestID}
	Update := bson.M{
		"$set": bson.M{
			"request_status": MongoConfig.RequestAccepted,
		},
	}

	_, err := GlobalVariables.MongoRequestsCollection.UpdateOne(context.Background(), filter, Update)
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

	_, err := GlobalVariables.MongoRequestsCollection.UpdateOne(context.Background(), filter, Update)
	if err != nil {
		log.Println(err)
	}
	DeleteRejectedRequestChat(RequestID)
}

func DeleteRejectedRequestChat(RequestID bson.ObjectID) {
	var request MongoConfig.Request
	filter := bson.M{"_id": RequestID}
	err := GlobalVariables.MongoRequestsCollection.FindOne(context.Background(), filter).Decode(&request)
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

	ChatErr := GlobalVariables.MongoChatsCollection.FindOne(context.Background(), chatFilter).Decode(&chat)

	if ChatErr != nil {
		log.Println(ChatErr)
	}

	messageFilter := bson.M{
		"chat_id": chat.ID,
	}

	_, MessagesErr := GlobalVariables.MongoMessagesCollection.DeleteMany(context.Background(), messageFilter)

	if MessagesErr != nil {
		log.Println(MessagesErr)
	}

	_, ChatDelErr := GlobalVariables.MongoChatsCollection.DeleteOne(context.Background(), bson.M{"_id": chat.ID})
	if ChatDelErr != nil {
		log.Println(ChatDelErr)
	}

	_, RequestDelErr := GlobalVariables.MongoRequestsCollection.DeleteOne(context.Background(), bson.M{"_id": RequestID})

	if RequestDelErr != nil {
		log.Println(RequestDelErr)
	}

}

func CacheMessageRedis(Message MongoConfig.Message) {
	// NOTE: If a message gets edited or deleted, the cache for the said conversation should be destroyed entireloy
	key := "chat:" + Message.ChatID.Hex() + ":messages"

	MarshaledMessage, err := json.Marshal(Message)
	if err != nil {
		log.Println(err)
		return
	}
	pipeline := Redis.Client.Pipeline()
	pipeline.LPush(context.Background(), key, MarshaledMessage)
	pipeline.LTrim(context.Background(), key, 0, 49)
	pipeline.Expire(context.Background(), key, 1*time.Hour)

	_, err = pipeline.Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func GetCachedMessages(ChatID bson.ObjectID, CurrentUserID bson.ObjectID) ([]MongoConfig.Message, error) {
	key := "chat:" + ChatID.Hex() + ":messages"
	msgs, err := Redis.Client.LRange(context.Background(), key, 0, 49).Result()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(msgs) == 0 {
		DBMessages, err := FetchLatestMessagesFromDB(ChatID)
		if err != nil {
			return nil, err
		}
		return DBMessages, nil
	}

	TargetMessages := make([]MongoConfig.Message, 0, len(msgs))
	for _, Message := range msgs {
		var TargetMessage MongoConfig.Message
		err := json.Unmarshal([]byte(Message), &TargetMessage)
		if err != nil {
			continue
		}
		if TargetMessage.SenderID == CurrentUserID {
			TargetMessage.IsCurrentUserSender = true
		} else {
			TargetMessage.IsCurrentUserSender = false
		}
		TargetMessages = append(TargetMessages, TargetMessage)
	}
	return TargetMessages, nil
}

func RetrieveAllUserChats(UserID bson.ObjectID) []MongoConfig.Chat {
	var TargetChats []MongoConfig.Chat
	UserChatsFilter := bson.M{
		"$or": []bson.M{
			{"creator_id": UserID},
			{"receiver_id": UserID},
		}}

	cursor, err := GlobalVariables.MongoChatsCollection.Find(context.Background(), UserChatsFilter)

	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &TargetChats)
	if err != nil {
		log.Println(err)
	}
	// Retrieving Messages
	for idx, TargetChat := range TargetChats {
		res, err := GetCachedMessages(TargetChat.ID, UserID)
		if err != nil {
			log.Println(err)
		}
		TargetChats[idx].Messages = res
		if TargetChat.CreatorID == UserID {
			DisplayedUserID := TargetChat.ReceiverID
			DisplayedUsername, err := GetUserData(DisplayedUserID)
			if err != nil {
				panic(err)
			} else {
				TargetChats[idx].DisplayedUsername = DisplayedUsername
			}
		} else {
			DisplayedUserID := TargetChat.CreatorID
			DisplayedUsername, err := GetUserData(DisplayedUserID)
			if err != nil {
				panic(err)
			} else {
				TargetChats[idx].DisplayedUsername = DisplayedUsername
			}
		}
	}

	return TargetChats
}

func FetchLatestMessagesFromDB(ChatID bson.ObjectID) ([]MongoConfig.Message, error) {
	// NOTE: A cursor will be needed for that purpose pointing to the latest message that the user has seen in the chat.
	filter := bson.M{"chat_id": ChatID}
	opts := options.Find().SetLimit(50)

	var TargetMessages []MongoConfig.Message

	cursor, err := GlobalVariables.MongoMessagesCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &TargetMessages)

	if err != nil {
		return nil, err
	}

	for _, Message := range TargetMessages {
		CacheMessageRedis(Message)
	}

	return TargetMessages, nil

}
