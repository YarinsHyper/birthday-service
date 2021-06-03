package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/yarinBenisty/birthday-service/manager"
	pb "github.com/yarinBenisty/birthday-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Service struct {
	Database           *mongo.Database
	PodcastsCollection *mongo.Collection
	m                  *manager.Manager
}

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
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	s.Database = client.Database("birthdayDB")
	s.PodcastsCollection = s.Database.Collection("birthdays")
	s.m = manager.NewManager(client.Database("birthdayDB"))

}

func (s *Service) CreateBirthday(ctx context.Context, req *pb.CreateBirthdayRequest) (*pb.BirthdayObject, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@0.0.0.0:27017"))
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Println(req)
	s.m = manager.NewManager(client.Database("birthdayDB"))
	b := bson.M{
		"name":           req.Name,
		"date":           req.Date,
		"personalNumber": req.PersonalNumber,
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, b)

	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("new birthday was created.\nbirthday id: ", insertResult.InsertedID)

	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}

func (s *Service) GetBirthday(ctx context.Context, req *pb.GetBirthdayRequest) (*pb.BirthdayObject, error) {
	fmt.Print("Read a birthday wroking succesfully.\n")
	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}

func (s *Service) UpdateBirthday(ctx context.Context, req *pb.UpdateBirthdayRequest) (*pb.BirthdayObject, error) {
	fmt.Print("Updated a birthday working succesfully.\n")
	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}

func (s *Service) DeleteBirthday(ctx context.Context, req *pb.DeleteBirthdayRequest) (*pb.BirthdayObject, error) {
	fmt.Print("Deleted a birthday working succesfully.\n")
	return &pb.BirthdayObject{Name: "yarin", Date: "35486715", PersonalNumber: "5615643214"}, nil
}
