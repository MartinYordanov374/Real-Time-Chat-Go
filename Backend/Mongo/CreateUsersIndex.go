package MongoConfig

import (
	"RealTimeChatApp/Backend/GlobalVariables"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateUsersIndex() error {
	_, err := GlobalVariables.MongoUsersCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "Username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return err
}
