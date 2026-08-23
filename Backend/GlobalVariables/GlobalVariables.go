package GlobalVariables

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"RealTimeChatApp/Backend/WebSockets"
)

var MongoClient *mongo.Client
var MongoUsersCollection *mongo.Collection
var MongoChatsCollection *mongo.Collection
var MongoMessagesCollection *mongo.Collection
var MongoRequestsCollection *mongo.Collection
var SessionDuration = 1*time.Hour
var CookieExpirationSeconds = 3600
var WebSocketsHub *WebSockets.Hub
