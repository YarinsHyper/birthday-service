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

const (
	birthdayDatabase        = "birthdayDB"
	birthdayCollection      = "birthdays"
	personalNumberParameter = "personalNumber"
	dateParameter           = "date"
	nameParameter           = "name"
)

// Service is a type of every method
type Service struct {
	Database           *mongo.Database
	BirthdayCollection *mongo.Collection
	pb.UnimplementedBirthdayFunctionsServer
}

// BirthdayModel is a model for coverting an object for type: birthdayObject
type BirthdayModel struct {
	Name           string
	Date           string
	PersonalNumber string
	ID             string `bson:"_id" json:"id,omitempty"`
}

// NewService is used each time i create a method
func NewService(url string) *Service {
	s := &Service{}
	// s.BirthdayCollection = mongoConnect.CNX.Database(birthdayDatabase).Collection(birthdayCollection)
	s.BirthdayCollection = mongoConnect.CNX.Database("birthdayDB").Collection("birthdays")

	return s
}

// asBirthdayObject is making any object type to: BirthdayObject type
func (bs *BirthdayModel) asBirthdayObject() *pb.BirthdayObject {
	birthdayO := &pb.BirthdayObject{
		Name:           bs.Name,
		Date:           bs.Date,
		PersonalNumber: bs.PersonalNumber,
		ID:             bs.ID,
	}
	return birthdayO
}

// CreateBirthday is creating a birthday object
func (s *Service) CreateBirthday(ctx context.Context, req *pb.CreateBirthdayRequest) (*pb.BirthdayObject, error) {

	b := bson.M{
		nameParameter:           req.Name,
		dateParameter:           req.Date,
		personalNumberParameter: req.PersonalNumber,
	}

	cursor, err := s.BirthdayCollection.InsertOne(ctx, b)
	if err != nil {
		log.Fatal(err)
	}

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{"_id": cursor.InsertedID})
	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)

	convertedBirthday := BirthdayModel.asBirthdayObject()

	return convertedBirthday, nil
}

// GetBirthday is getting a birthday object by personalNumber parameter
func (s *Service) GetBirthday(ctx context.Context, req *pb.GetBirthdayRequest) (*pb.BirthdayObject, error) {

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{personalNumberParameter: req.PersonalNumber})
	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	return convertedBirthday, nil
}

// GetAllBirthdays is creating all birthday objects
func (s *Service) GetAllBirthdays(ctx context.Context, req *pb.GetAllBirthdaysRequest) (*pb.GetAllBirthdaysResponse, error) {

	cursor, err := s.BirthdayCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var mongoBirthdays []BirthdayModel
	if err = cursor.All(ctx, &mongoBirthdays); err != nil {
		log.Fatal(err)
	}
	birthdaysArray := make([]*pb.BirthdayObject, 0, len(mongoBirthdays))

	// This for loops inside the mongoBirthdays, retrieves each one
	// and appends it to a new birthdaysArray by a new type.
	for index := 0; index < len(mongoBirthdays); index++ {

		birthday := mongoBirthdays[index]
		convertedBirthday := birthday.asBirthdayObject()
		birthdaysArray = append(birthdaysArray, convertedBirthday)
	}

	return &pb.GetAllBirthdaysResponse{Birthdays: birthdaysArray}, nil
}

// UpdateBirthday is updating a birthday object
func (s *Service) UpdateBirthday(ctx context.Context, req *pb.UpdateBirthdayRequest) (*pb.BirthdayObject, error) {

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: personalNumberParameter, Value: req.PersonalNumber}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: nameParameter, Value: req.Name}, {Key: nameParameter, Value: req.Date}, {Key: personalNumberParameter, Value: req.PersonalNumber}}}}
	result, err := s.BirthdayCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{personalNumberParameter: req.PersonalNumber})

	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	fmt.Printf("updated %v Document/s!\n", result.ModifiedCount)

	return convertedBirthday, nil
}

// DeleteBirthday is deleting a birthday object
func (s *Service) DeleteBirthday(ctx context.Context, req *pb.DeleteBirthdayRequest) (*pb.DeleteBirthdayResponse, error) {

	result, err := s.BirthdayCollection.DeleteOne(ctx, bson.M{personalNumberParameter: req.PersonalNumber})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v birthday(s)\n", result.DeletedCount)

	return &pb.DeleteBirthdayResponse{}, nil
}
