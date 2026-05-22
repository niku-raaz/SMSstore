package kafka

import (
	"SMSstore/models"
	"SMSstore/service"
	"context"
	"encoding/json"
	"log"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

func StartConsumer() {

	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "sms-topic",
		GroupID: "sms-store-group",

		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	log.Println("Kafka Consumer Started")

	for {

		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("Kafka Read Error:", err)
			continue
		}

		log.Println("Message Received: ", string(msg.Value))

		var sms models.SMS

		err = json.Unmarshal(msg.Value, &sms)

		if err != nil {
			log.Println("JSON Parse Error:", err)
			continue
		}

		sms.CreatedAt = time.Now()

		err = service.SaveSMS(sms)

		if err != nil {
			log.Println("Mongo Save Error:", err)
			continue
		}

		log.Println("SMS Saved Successfully")
	}
}
