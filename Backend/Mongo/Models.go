package MongoConfig
import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID	bson.ObjectID `json:"ID" bson:"_id,omitempty"`
	Username		string		`json:"Username"`
	Password	string		`json:"Password"`
	Email			string		`json:"Email"`
}

type Test struct{
	SessionID string
}
