package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuario struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Surname string             `bson:"surname" json:"surname"`
	Email   string             `bson:"email" json:"email"`
}
