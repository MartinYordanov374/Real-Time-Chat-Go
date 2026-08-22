package MongoConfig

import(
	"log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"RealTimeChatApp/Backend/GlobalVariables"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToMongo() (*mongo.Client, error){
	uri := "mongodb://mongo:27017/RTC"
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil{
		log.Println(err)
	}else{
		log.Println("Connected to DB successfully!")
		GlobalVariables.MongoClient = client
		UsersCollection := GlobalVariables.MongoClient.Database("RTC").Collection("Users")
		ChatsCollection := GlobalVariables.MongoClient.Database("RTC").Collection("Chats")

		GlobalVariables.MongoUsersCollection = UsersCollection
		GlobalVariables.MongoChatsCollection = ChatsCollection

	 }

	return client, nil;
}
