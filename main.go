package main

import (
	"github.com/joho/godotenv"
	"os"

	"SMSstore/config"
)

func main() {

	godotenv.Load()

	uri := os.Getenv("MONGO_URI")

	config.ConnectMongo(uri)

	// start HTTP server here
}
