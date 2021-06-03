package manager

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/yarinBenisty/birthday-service/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type Manager struct{
	c *mongo.Database;
	repository *repository.Repository;
}

func NewManager(c *mongo.Database) *Manager{
	return &Manager{c:c, repository: repository.NewRepository(c)}
}

func (m *Manager)CreateBirthdayManager(personObject bson.M){
	m.repository.CreateBirthdayRepository(personObject)
}
