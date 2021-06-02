package controller

import (
	"context"
	"fmt"
	"log"

	pb "github.com/yarinBenisty/birthday-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Service struct{}

// NewService creates a Service do mongo stuff and returns it
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
	ctx := context.Background() // context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("birthdayDB")
	podcastsCollection := database.Collection("birthdays")

	b := bson.M{
		"name":           "benny",
		"date":           162068714564501,
		"personalNumber": "246978284",
	}

	insertResult, err := podcastsCollection.InsertOne(ctx, b)

	if err != nil {
		panic(err)
	}
	fmt.Println("new birthday was created.\nbirthday id: ", insertResult.InsertedID)
}

func (s *Service) CreateBirthday(ctx context.Context, req *pb.CreateBirthdayRequest) (*pb.BirthdayObject, error) {
	fmt.Print("Created a birthday succesfully.\n")
	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}

func (s *Service) GetBirthday(ctx context.Context, req *pb.GetBirthdayRequest) (*pb.BirthdayObject, error) {
	fmt.Print("Read a birthday succesfully.\n")
	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}

func (s *Service) UpdateBirthday(ctx context.Context, req *pb.UpdateBirthdayRequest) (*pb.BirthdayObject, error) {
	fmt.Print("Updated a birthday succesfully.\n")
	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}

func (s *Service) DeleteBirthday(ctx context.Context, req *pb.DeleteBirthdayRequest) (*pb.BirthdayObject, error) {
	fmt.Print("Deleted a birthday succesfully.\n")
	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}
