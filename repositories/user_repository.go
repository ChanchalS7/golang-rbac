package repositories

import (
	"context"

	"github.com/ChanchalS7/golang-rbac/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/primitive"
)

//UseRepository handles the CRUD operations with the MongoDB database

type UserRepository struct {

	Collection *mongo.Collection
}
//CreateUse insert a new user into database

func (repo *UserRepository) CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	return repo.Collection.InsertOne(context.TODO(),user)
}

//GetUserByID fetches a user by their ID
func (repo *UserRepository) GetUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := repo.Collection.FindOne(context.TODO(), bson.M{"_id":id}).Decode(&user)
	return user, err
}
//UpdateUser updates a user's details in the database
func (repo *UserRepository) UpdateUser(id primitive.ObjectID, user models.User) (*mongo.UpdateResult, error) {
return repo.Collection.UpdateOne(context.TODO(), bson.M{"_id":id}, bson.M{"$set":user})

}

//DeleteUser remove a user from the database by their ID
func (repo *UserRepository) DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return repo.Collection.DeleteOne(context.TODO(), bson.M{"_id":id})
}