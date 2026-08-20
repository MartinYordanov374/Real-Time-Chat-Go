package Redis
import "go.mongodb.org/mongo-driver/v2/bson"

type Session struct{
	SessionID string
	UserID bson.ObjectID `json:"ID" bson: "_id"`
}
