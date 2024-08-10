package services

import (
	"github.com/ChanchalS7/golang-rbac/models"
	"github.com/ChanchalS7/golang-rbac/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repo *repositories.UserRepository
}

//CreateUser create a nwe user
func (service *UserService) CreateUser(user models.User) (*models.User, error){
	_,err:= service.Repo.CreateUser(user)
	return &user,err
}

//GEtUserByID returns a user by thier ID
func (service *UserService) GetUserByID(id string) (*models.User, error){
	objectID , _ := primitive.ObjectIDFromHex(id)
	user, err := service.Repo.GetUserByID(objectID)
	return &user, err
}

//UpdateUser udpates a user's details
func (service *UserService) UpdateUser(id string, user models.User) (*models.User, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err:= service.Repo.UpdateUser(objectID, user)
	return &user, err 
}

//DeleteUser deletes a user by thier ID
func (service *UserService) DeleteUser(id string) error{
	objectID, _:= primitive.ObjectIDFromHex(id)
	_,err:= service.Repo.DeleteUser(objectID)
return err

}