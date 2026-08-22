package Redis

import "go.mongodb.org/mongo-driver/v2/bson"

type Session struct{
	UserID bson.ObjectID `bson: UserID`
}
