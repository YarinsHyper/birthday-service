package mongomanager

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CNX is a mongo client
var CNX = Connection()

//Connection is a function that creates a mongo client and
//connects it to the db and pings/checks the connection
func Connection() *mongo.Client {

	// loading the dotenv parameters
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))

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
