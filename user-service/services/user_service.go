package services

import (
	"context"
	"log"

	"time"

	"github.com/wjahatsyed/user-service/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func InitMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	UserCollection = client.Database("foodapp").Collection("users")
	log.Println("✅ MongoDB Connected!")
}

func InsertUser(ctx context.Context, user models.User) (interface{}, error) {
	result, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
