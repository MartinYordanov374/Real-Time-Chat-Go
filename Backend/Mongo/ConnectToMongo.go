package MongoConfig

import (
	"RealTimeChatApp/Backend/GlobalVariables"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToMongo() (*mongo.Client, error) {
	uri := "mongodb://mongo:27017/RTC"
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to DB successfully!")
		GlobalVariables.MongoClient = client
		UsersCollection := GlobalVariables.MongoClient.Database("RTC").Collection("Users")
		ChatsCollection := GlobalVariables.MongoClient.Database("RTC").Collection("Chats")
		MessagesCollection := GlobalVariables.MongoClient.Database("RTC").Collection("Messages")
		RequestsCollection := GlobalVariables.MongoClient.Database("RTC").Collection("Requests")
		GlobalVariables.MongoUsersCollection = UsersCollection
		GlobalVariables.MongoChatsCollection = ChatsCollection
		GlobalVariables.MongoMessagesCollection = MessagesCollection
		GlobalVariables.MongoRequestsCollection = RequestsCollection

		err := CreateUsersIndex()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return client, nil
}
