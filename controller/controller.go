package controller

import (
	"context"

	pb "github.com/yarinBenisty/birthday-service/proto"
)

// Controller is an interface for the business logic of the birthday.Service
type Controller interface {
	CreateBirthday(ctx context.Context, Name string, Date string, PersonalNumber string) (*pb.BirthdayObject, error)
	GetBirthday(ctx context.Context, PersonalNumber string) (*pb.BirthdayObject, error)
	GetAllBirthdays(ctx context.Context) (*pb.GetAllBirthdaysResponse, error)
	DeleteBirthday(ctx context.Context, PersonalNumber string) (*pb.DeleteBirthdayResponse, error)
}