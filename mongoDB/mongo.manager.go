package mongomanager

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CNX is a mongo client
var CNX = Connection()

//Connection is a function that creates a mongo client and
//connects it to the db and pings/checks the connection
func Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@0.0.0.0:27017")

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to connect to mongo!")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to connect to mongo!")
	}

	return client
}
