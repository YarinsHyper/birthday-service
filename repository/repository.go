package repository

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct{
	c *mongo.Database;
}

func NewRepository(c *mongo.Database) *Repository{
	return &Repository{c:c}
}

func (r *Repository)CreateBirthdayRepository(personObject bson.M){
	collection := r.c.Collection("birthdays")
	insertResult, err := collection.InsertOne(context.Background(), personObject)

	if err != nil {	
		fmt.Println("error: ", err)
	}
	fmt.Println("new birthday was created.\nbirthday id: ", insertResult.InsertedID)

}
