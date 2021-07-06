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
	// BirthdayDatabase DB name
	BirthdayDatabase = "birthdayDB"
	// BirthdayCollection collection name
	BirthdayCollection = "birthdays"
	// PersonalNumberParameter object parameter
	PersonalNumberParameter = "personalNumber"
	// DateParameter object parameter
	DateParameter = "date"
	// NameParameter object parameter
	NameParameter = "name"
)

// Service is a type that every method uses to get
// all the info needed in order to work properly
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

// NewService creaes a new service and returns is
func NewService(url string) *Service {
	s := &Service{}
	s.BirthdayCollection = mongoConnect.CNX.Database(BirthdayDatabase).Collection(BirthdayCollection)

	return s
}

// asBirthdayObject converts BirthdayModel type to: BirthdayObject type
func (bs *BirthdayModel) asBirthdayObject() *pb.BirthdayObject {
	birthdayO := &pb.BirthdayObject{
		Name:           bs.Name,
		Date:           bs.Date,
		PersonalNumber: bs.PersonalNumber,
		ID:             bs.ID,
	}
	return birthdayO
}

// CreateBirthday is UPSERTING which means: if the given object
// exists its being overridden, if it doesn't its inserting the object
func (s *Service) CreateBirthday(ctx context.Context, req *pb.CreateBirthdayRequest) (*pb.BirthdayObject, error) {

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: PersonalNumberParameter, Value: req.PersonalNumber}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: NameParameter, Value: req.Name}, {Key: DateParameter, Value: req.Date}, {Key: PersonalNumberParameter, Value: req.PersonalNumber}}}}
	result, err := s.BirthdayCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{PersonalNumberParameter: req.PersonalNumber})
	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	fmt.Printf("changed %v Document/s!\n", result.ModifiedCount)

	return convertedBirthday, nil
}

// GetBirthday returns a birthday object by personalNumber
func (s *Service) GetBirthday(ctx context.Context, req *pb.GetBirthdayRequest) (*pb.BirthdayObject, error) {

	birthday := s.BirthdayCollection.FindOne(ctx, bson.M{PersonalNumberParameter: req.PersonalNumber})
	BirthdayModel := &BirthdayModel{}
	birthday.Decode(BirthdayModel)
	convertedBirthday := BirthdayModel.asBirthdayObject()

	return convertedBirthday, nil
}

// GetAllBirthdays returns all birthday objects in the DB
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

// DeleteBirthday deletes a birthday object by it's personalNumber
func (s *Service) DeleteBirthday(ctx context.Context, req *pb.DeleteBirthdayRequest) (*pb.DeleteBirthdayResponse, error) {

	result, err := s.BirthdayCollection.DeleteOne(ctx, bson.M{PersonalNumberParameter: req.PersonalNumber})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("removed %v document/s\n", result.DeletedCount)

	return &pb.DeleteBirthdayResponse{}, nil
}
