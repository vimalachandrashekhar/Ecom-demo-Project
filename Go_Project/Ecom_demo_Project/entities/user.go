package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `bson:"_id"`
	FirstName       string             `json:"firstName" bson:"firstName,required"`
	LastName        string             `json:"lastName" bson:"lastName,required"`
	Age             int                `json:"age" bson:"age,required"`
	Email           string             `json:"email" bson:"email,required"`
	Password        string             `json:"password" bson:"password,required"`
	ConfirmPassword string             `json:"confirmPassword" bson:"confirmPassword,required"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserResponse struct {
	Response string `json:"response"`
}

type Login struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required,min=8"`
}
