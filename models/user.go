package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//user represents a user in database
type User struct {
	ID 		 primitive.ObjectID	`bson:"_id, omitEmpty"`
	Name	 string 				`bson:"name,omitEmpty"`
	Email	 string					`bson:"email,omitempty"`
	Password string				`bson:"password,omitempty"`

}