package MongoConfig

import(
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

func ConnectToMongo(){
	uri := "mongodb://mongo:27017/RTC"
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil{
		log.Println(err)
	}else{
		log.Println("Connected to db successfully")
		log.Println(client)
	}
}
