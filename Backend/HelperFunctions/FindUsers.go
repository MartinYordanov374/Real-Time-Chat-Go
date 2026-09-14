package HelperFunctions
import (
	"RealTimeChatApp/Backend/GlobalVariables"
	MongoConfig "RealTimeChatApp/Backend/Mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
func FindUsers(TargetUsername string) ([]MongoConfig.User, error){
var TargetUsers []MongoConfig.User
	filter := bson.M{
		"username": bson.M{
			"$gte": TargetUsername,
			"$lt":  TargetUsername + "\uffff",
		},
	}
	opts := options.Find().SetLimit(5)
	// NOTE: This functionality should first look for the user's CONTACTS(chats with users) and then look for other users.
	// If a user is found and they do not have a chat with them, the user's username and ID should be displayed alone in the
	// contacts list (a username should be unique). The user is given the option to send a message to the said person.
	// After the message has been sent a box appears that asks the user to wait for the other user to accept the chat request.
	// TODO: Rewrite this to match usernames by regex or soemthing once you get the functionality working for a specific Username
	// TODO: Do not return the entire user object, just the information necessary
	cursor, err := GlobalVariables.MongoUsersCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return []MongoConfig.User{}, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &TargetUsers)
	if err != nil {
		return []MongoConfig.User{}, err
	}

	return TargetUsers, nil
}
