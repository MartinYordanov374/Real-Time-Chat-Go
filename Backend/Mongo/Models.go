package MongoConfig

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	Email    string        `json:"email" bson:"email"`
}

type Chat struct {
	ID           bson.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatorID    bson.ObjectID `json:"creatorId" bson:"creator_id"`
	ReceiverID   bson.ObjectID `json:"receiverId" bson:"receiver_id"`
	CreationDate time.Time     `json:"creationDate" bson:"creation_date"`
}

type Message struct {
	ID          bson.ObjectID `json:"id" bson:"_id,omitempty"`
	ChatID      bson.ObjectID `json:"chatId" bson:"chat_id"`
	TextContent string        `json:"textContent" bson:"text_content"`
	TimeStamp   time.Time     `json:"timeStamp" bson:"time_stamp"`
	SenderID    bson.ObjectID `json:"senderId" bson:"sender_id"`
}

type Request struct {
	ID         bson.ObjectID `json:"id" bson:"_id,omitempty"`
	SenderID   bson.ObjectID `json:"senderId" bson:"sender_id"`
	ReceiverID bson.ObjectID `json:"receiverId" bson:"receiver_id"`
	Status     RequestStatus `json:"requestStatus" bson:"request_status"`
	TimeStamp  time.Time     `json:"timeStamp" bson:"time_stamp"`
}

type Inbox struct {
	ID         bson.ObjectID `json:"id" bson:"_id,omitempty"`
	SenderID   bson.ObjectID `json:"senderId" bson:"sender_id"`
	ReceiverID bson.ObjectID `json:"receiverId" bson:"receiver_id"`
	Category   bool          `json:"category" bson:"category"`
	TimeStamp  time.Time     `json:"timeStamp" bson:"time_stamp"`
}
