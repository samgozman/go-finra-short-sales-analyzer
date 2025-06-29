package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This method closes mongoDB connection and   context.
func Close(ctx context.Context, client *mongo.Client) {

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// Connect to the mongo database
func Connect(ctx context.Context, uri string, auth options.Credential) (*mongo.Client, error) {

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(auth))
	return client, err
}
