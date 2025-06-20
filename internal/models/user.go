package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name     string             `bson:"name" json:"name"`
    Email    string             `json:"email"`
    Password string             `json:"password"`
    Role     string             `bson:"role" json:"role"` // user | admin
}
