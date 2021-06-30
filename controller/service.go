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

//Service is a type of every method
type Service struct {
	Database           *mongo.Database
	BirthdayCollection *mongo.Collection
	pb.UnimplementedBirthdayFunctionsServer
}

//BirthdayModel is a model for coverting an object for type: birthdayObject
type BirthdayModel struct {
	Name           string
	Date           string
	PersonalNumber string
	ID             string `bson:"_id" json:"id,omitempty"`
}

//NewService is used each time i create a method
func NewService(url string) *Service {
	s := &Service{}
	s.BirthdayCollection = mongoConnect.CNX.Database("birthdayDB").Collection("birthdays")

	return s
}

func (bs *BirthdayModel) asBirthdayObject() *pb.BirthdayObject {
	birthdayO := &pb.BirthdayObject{
		Name:           bs.Name,
		Date:           bs.Date,
		PersonalNumber: bs.PersonalNumber,
		ID:             bs.ID,
	}
	return birthdayO
}

//CreateBirthday is creating a birthday object
func (s *Service) CreateBirthday(ctx context.Context, req *pb.CreateBirthdayRequest) (*pb.BirthdayObject, error) {

	b := bson.M{
		"name":           req.Name,
		"date":           req.Date,
		"personalNumber": req.PersonalNumber,
	}

	cursor, err := s.BirthdayCollection.InsertOne(ctx, b)
	if err != nil {
		fmt.Println("error: ", err)
	}

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{"_id": cursor.InsertedID})
	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)

	convertedBirthday := BirthdayModel.asBirthdayObject()
	fmt.Println(convertedBirthday)

	return convertedBirthday, nil
}

//GetBirthday is getting a birthday object
func (s *Service) GetBirthday(ctx context.Context, req *pb.GetBirthdayRequest) (*pb.BirthdayObject, error) {

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{"personalNumber": req.PersonalNumber})
	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	return convertedBirthday, nil
}

//GetAllBirthday is creating all birthday objects
func (s *Service) GetAllBirthday(ctx context.Context, req *pb.GetAllBirthdayRequest) (*pb.GetAllBirthdayResponse, error) {

	cursor, err := s.BirthdayCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var mongoBirthdays []BirthdayModel
	if err = cursor.All(ctx, &mongoBirthdays); err != nil {
		log.Fatal(err)
	}
	birthdaysArray := make([]*pb.BirthdayObject, 0, len(mongoBirthdays))

	for index := 0; index < len(mongoBirthdays); index++ {

		birthday := mongoBirthdays[index]
		convertedBirthday := birthday.asBirthdayObject()
		birthdaysArray = append(birthdaysArray, convertedBirthday)

		fmt.Println(convertedBirthday)
	}

	return &pb.GetAllBirthdayResponse{Birthdays: birthdaysArray}, nil
}

// UpdateBirthday is updating a birthday object
func (s *Service) UpdateBirthday(ctx context.Context, req *pb.UpdateBirthdayRequest) (*pb.BirthdayObject, error) {

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "personalNumber", Value: req.PersonalNumber}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: req.Name}, {Key: "date", Value: req.Date}, {Key: "personalNumber", Value: req.PersonalNumber}}}}
	result, err := s.BirthdayCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{"personalNumber": req.PersonalNumber})

	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	fmt.Printf("updated %v Document/s!\n", result.ModifiedCount)

	return convertedBirthday, nil
}

// DeleteBirthday is deleting a birthday object
func (s *Service) DeleteBirthday(ctx context.Context, req *pb.DeleteBirthdayRequest) (*pb.DeleteBirthdayResponse, error) {

	result, err := s.BirthdayCollection.DeleteOne(ctx, bson.M{"personalNumber": req.PersonalNumber})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v birthday(s)\n", result.DeletedCount)

	return &pb.DeleteBirthdayResponse{}, nil
}
