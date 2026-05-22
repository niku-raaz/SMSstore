package repository

import (
	"SMSstore/config"
	"SMSstore/models"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func SaveSMS(sms models.SMS) error {

	_, err := config.SMSCollection.InsertOne(
		context.Background(),
		sms,
	)

	return err
}

func GetMessagesByUserID(userID int64) ([]models.SMS, error) {

	filter := bson.M{
		"userId": userID,
	}

	cursor, err := config.SMSCollection.Find(
		context.Background(),
		filter,
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var messages []models.SMS

	if err := cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
