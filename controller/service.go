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

type BirthdayStruct struct {
	Name           string
	Date           string
	PersonalNumber string
	ID             string `bson:"_id" json:"id,omitempty"`
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

}

func (bs *BirthdayStruct) asBirthdayObject() *pb.BirthdayObject {
	birthdayO := &pb.BirthdayObject{
		Name:           bs.Name,
		Date:           bs.Date,
		PersonalNumber: bs.PersonalNumber,
		ID:             bs.ID,
	}
	return birthdayO
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
	birthdayCollection := database.Collection("birthdays")

	b := bson.M{
		"name":           req.Name,
		"date":           req.Date,
		"personalNumber": req.PersonalNumber,
	}

	cursor, err := birthdayCollection.InsertOne(ctx, b)
	if err != nil {
		fmt.Println("error: ", err)
	}

	birthday := birthdayCollection.FindOne(ctx, bson.M{"_id": cursor.InsertedID})
	BirthdayModel := &BirthdayStruct{}
	birthday.Decode(BirthdayModel)

	convertedBirthday := BirthdayModel.asBirthdayObject()
	fmt.Println(convertedBirthday)

	return convertedBirthday, nil
}

func (s *Service) GetBirthday(ctx context.Context, req *pb.GetBirthdayRequest) (*pb.BirthdayObject, error) {
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
	birthdaysCollection := database.Collection("birthdays")

	cursor, err := birthdaysCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("error: ", err)
	}

	var birthdays []bson.M
	if err = cursor.All(ctx, &birthdays); err != nil {
		log.Fatal(err)
	}

	for _, birthdays := range birthdays {
		fmt.Println(birthdays["name"], birthdays["date"], birthdays["personalNumber"], birthdays["_id"])
	}
	fmt.Println(birthdays)
	// birthdayModel := &BirthdayStruct{}
	// birthdays.Decode(birthdayModel)

	// birthdays := &BirthdayStruct{}
	// convertedBirthdays := birthdays.asBirthdayObject()

	return &pb.BirthdayObject{}, nil
}

func (s *Service) UpdateBirthday(ctx context.Context, req *pb.UpdateBirthdayRequest) (*pb.BirthdayObject, error) {
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
	birthdaysCollection := database.Collection("birthdays")
	s.m = manager.NewManager(client.Database("birthdayDB"))

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"personalNumber", req.PersonalNumber}}
	update := bson.D{{"$set", bson.D{{"name", "yarin"}}}}
	result, err := birthdaysCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("updated %v Documents!\n", result.ModifiedCount)

	return &pb.BirthdayObject{}, nil
}

func (s *Service) DeleteBirthday(ctx context.Context, req *pb.DeleteBirthdayRequest) (*pb.DeleteBirthdayResponse, error) {
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
	birthdaysCollection := database.Collection("birthdays")
	s.m = manager.NewManager(client.Database("birthdayDB"))

	result, err := birthdaysCollection.DeleteOne(ctx, bson.M{"name": "shalev-goldshtein"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v birthday(s)\n", result.DeletedCount)

	return &pb.DeleteBirthdayResponse{}, nil
}
