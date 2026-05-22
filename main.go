package main

import (
	"SMSstore/handlers"
	"SMSstore/kafka"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"SMSstore/config"
)

func main() {

	godotenv.Load()

	uri := os.Getenv("MONGO_URI")

	config.ConnectMongo(uri)
	go kafka.StartConsumer()

	// start HTTP server here
	//http.HandleFunc("/test", handlers.TestInsert)

	http.HandleFunc(
		"GET /v1/user/{userId}/messages",
		handlers.GetMessagesByUser,
	)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
