package storage

import (
	"math/rand/v2"
	"strconv"
	"users/internal/models"
)

type Storage interface {
	GetUser(id uint) models.User
	CreateUser(name, email string) (id string)
}

func NewStorage() Storage {
	return &MockStorage{
		users: make(map[uint]models.User),
	}
}

type MockStorage struct {
	users map[uint]models.User
}

func (s *MockStorage) GetUser(id uint) models.User {
	return s.users[id]
}

func (s *MockStorage) CreateUser(name, email string) string {
	id := rand.Int()
	if id < 0 {
		id = -id
	}
	s.users[uint(id)] = models.User{Name: name, Email: email}
	return strconv.Itoa(int(id))
}
