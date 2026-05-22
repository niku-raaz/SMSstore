SMS Store Service

SMS Store Service is a Go microservice responsible for consuming SMS events from Kafka, storing them in MongoDB, and providing APIs to retrieve SMS history for users.

This service acts as the persistence layer of the SMS platform and receives SMS events asynchronously from the SMS Sender Service through Kafka.

Basic Setup to Run the Project
1. Start the Springboot project and all its required services.
2. Ensure MongoDB Atlas is configured and the connection URI is present in the .env file.
3. Navigate to the project directory.
4. Run the Go application: go run main.go




Tech Stack
Go
MongoDB Atlas
Apache Kafka
net/http
Author

Niku Raj

Spring Boot | Distributed Systems | Backend Development | Kafka | Redis | Go | MongoDB
