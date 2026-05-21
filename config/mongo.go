package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

var SMSCollection *mongo.Collection

func ConnectMongo(uri string) {
	//fmt.Printf("URI = [%s]\n", uri)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		fmt.Println("Error Connecting to MongoDB")
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB Ping Failed:", err)
	}

	log.Println("Connected to MongoDB Atlas")

	db := client.Database("sms-store")

	SMSCollection = db.Collection("messages")
}
