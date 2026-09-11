package GlobalVariables

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var MongoClient *mongo.Client
var MongoUsersCollection *mongo.Collection
var MongoChatsCollection *mongo.Collection
var MongoMessagesCollection *mongo.Collection
var MongoRequestsCollection *mongo.Collection
var SessionDuration = 1*time.Hour
var CookieExpirationSeconds = 3600
type RedisMessage struct {
	ReceiverID bson.ObjectID `json:"ReceiverID"`
	SenderID bson.ObjectID `json:"SenderID"`
	Content string		 `json:"Content"`
}
