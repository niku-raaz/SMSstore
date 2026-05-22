package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type SMS struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    int64         `bson:"userId" json:"userId"`
	Recipient string        `bson:"recipient" json:"recipient"`
	Message   string        `bson:"message" json:"message"`
	Status    string        `bson:"status" json:"status"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}
