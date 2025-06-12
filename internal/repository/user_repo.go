package repository

import (
    "context"
    "store/config"
    "store/internal/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func UserCollection() *mongo.Collection {
    return config.DB.Collection("users")
}

func FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
    var user models.User
    err := UserCollection().FindOne(ctx, bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func CreateUser(ctx context.Context, user *models.User) error {
    _, err := UserCollection().InsertOne(ctx, user)
    return err
}
