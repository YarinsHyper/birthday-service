package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Service struct {
}

// NewService creates a Service and returns it.
func NewService(url string) *Service {
	s := &Service{}
	s.mongoConnection(url)
	return s
}

func (s *Service) mongoConnection(url string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	database := client.Database("birthdayDB")
	podcastsCollection := database.Collection("birthdays")

	b := bson.M{
		"name":           "yarin benisty",
		"date":           21321454565463,
		"personalNumber": "208374635",
	}

	insertResult, err := podcastsCollection.InsertOne(ctx, b)
	if err != nil {
		panic(err)
	}
	fmt.Println("new birthday was created.\nbirthday id: ", insertResult.InsertedID)
}
