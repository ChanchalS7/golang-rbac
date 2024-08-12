
package tests

import (
	"errors"
	"testing"

	"github.com/ChanchalS7/golang-rbac/models"
	"github.com/ChanchalS7/golang-rbac/services"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockRepo is a mock implementation of UserRepoInterface for testing
type MockRepo struct {
	users map[primitive.ObjectID]models.User
}

func (repo *MockRepo) CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	repo.users[user.ID] = user
	return &mongo.InsertOneResult{}, nil
}

func (repo *MockRepo) GetUserByID(id primitive.ObjectID) (models.User, error) {
	user, exists := repo.users[id]
	if !exists {
		return models.User{}, errors.New("User not found")
	}
	return user, nil
}
func TestCreateUser(t *testing.T) {
	mockRepo := &MockRepo{users: make(map[primitive.ObjectID]models.User)}
	userService := &services.UserService{Repo: mockRepo}

	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	createdUser, err := userService.CreateUser(user)

	assert.Nil(t, err)
	assert.Equal(t, "Test User", createdUser.Name)
}
