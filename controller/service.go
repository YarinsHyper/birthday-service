package controller

import (
	"context"
	"fmt"
	"log"

	mongoConnect "github.com/yarinBenisty/birthday-service/mongoDB"
	pb "github.com/yarinBenisty/birthday-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	Database           *mongo.Database
	BirthdayCollection *mongo.Collection
}

type BirthdayStruct struct {
	Name           string
	Date           string
	PersonalNumber string
	ID             string `bson:"_id" json:"id,omitempty"`
}

func NewService(url string) *Service {
	s := &Service{}
	return s
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
	birthdayCollection := mongoConnect.CNX.Database("birthdayDB").Collection("birthdays")

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
	birthdayCollection := mongoConnect.CNX.Database("birthdayDB").Collection("birthdays")

	birthday := birthdayCollection.FindOne(ctx, bson.M{"personalNumber": req.PersonalNumber})
	BirthdayModel := &BirthdayStruct{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	return convertedBirthday, nil
}

func (s *Service) GetAllBirthday(ctx context.Context, req *pb.GetAllBirthdayRequest) (*pb.BirthdayObject, error) {
	birthdayCollection := mongoConnect.CNX.Database("birthdayDB").Collection("birthdays")

	birthday := birthdayCollection.FindOne(ctx, bson.M{})
	BirthdayModel := &BirthdayStruct{}
	birthday.Decode(BirthdayModel)
	convertedBirthdays := BirthdayModel.asBirthdayObject()

	return convertedBirthdays, nil
}

func (s *Service) UpdateBirthday(ctx context.Context, req *pb.UpdateBirthdayRequest) (*pb.BirthdayObject, error) {
	birthdayCollection := mongoConnect.CNX.Database("birthdayDB").Collection("birthdays")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"personalNumber", req.PersonalNumber}}
	update := bson.D{{"$set", bson.D{{"name", req.Name}, {"date", req.Date}, {"personalNumber", req.PersonalNumber}}}}
	result, err := birthdayCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	birthday := birthdayCollection.FindOne(ctx, bson.M{"personalNumber": req.PersonalNumber})

	BirthdayModel := &BirthdayStruct{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	fmt.Printf("updated %v Document/s!\n", result.ModifiedCount)

	return convertedBirthday, nil
}

func (s *Service) DeleteBirthday(ctx context.Context, req *pb.DeleteBirthdayRequest) (*pb.DeleteBirthdayResponse, error) {
	birthdayCollection := mongoConnect.CNX.Database("birthdayDB").Collection("birthdays")

	result, err := birthdayCollection.DeleteOne(ctx, bson.M{"personalNumber": req.PersonalNumber})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v birthday(s)\n", result.DeletedCount)

	return &pb.DeleteBirthdayResponse{}, nil
}
