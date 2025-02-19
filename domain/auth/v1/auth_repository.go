package auth

import (
    "context"
    "your_project/config"
    "your_project/common/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
    collection *mongo.Collection
}

func NewAuthRepository() *authRepository {
    return &authRepository{
        collection: config.DB.Collection("users"),
    }
}

func (repo *authRepository) FindUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := repo.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (repo *authRepository) SaveUser(user *models.User) (*mongo.InsertOneResult, error) {
    return repo.collection.InsertOne(context.TODO(), user)
}
