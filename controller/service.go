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

type Service struct {
	Database           *mongo.Database
	PodcastsCollection *mongo.Collection
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
	birthday := birthdaysCollection.FindOne(ctx, bson.M{"personalNumber": req.PersonalNumber})

	BirthdayModel := &BirthdayStruct{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	return convertedBirthday, nil
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

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"personalNumber", req.PersonalNumber}}
	update := bson.D{{"$set", bson.D{{"name", req.Name}, {"date", req.Date}, {"personalNumber", req.PersonalNumber}}}}
	result, err := birthdaysCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	birthday := birthdaysCollection.FindOne(ctx, bson.M{"personalNumber": req.PersonalNumber})

	BirthdayModel := &BirthdayStruct{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	fmt.Println(convertedBirthday)

	fmt.Printf("updated %v Document/s!\n", result.ModifiedCount)

	return convertedBirthday, nil
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

	result, err := birthdaysCollection.DeleteOne(ctx, bson.M{"personalNumber": req.PersonalNumber})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v birthday(s)\n", result.DeletedCount)

	return &pb.DeleteBirthdayResponse{}, nil
}
