# Polyglot Distributed SMS Service

This repository contains the solution for refactoring a monolithic SMS Notification Service into two separate, communicating microservices. 

### Project Overview
* **SMS Sender (`SMSSender`)**: A Java Spring Boot application acting as a Kafka Event Producer. It functions as a gateway, checking a Redis blocklist before sending messages to a mocked 3rd-party vendor.
* **SMS Store (`SMSstore`)**: A GoLang application acting as a Kafka Event Consumer and standard HTTP server. It consumes events and persists the messages in MongoDB.

---

## Architecture & Data Flow

1. A client sends a request to the **Java SMS Sender** API.
2. The Java service checks **Redis** to ensure the user is not in the blocklist.
3. The Java service calls a mocked 3rd-party vendor API to send the SMS (mocked with Success/Fail).
4. The Java service pushes the SMS event (with status) to **Kafka**.
5. The **GoLang SMS Store** listens to the Kafka topic, consumes the event, and stores the record in **MongoDB**.
6. Clients can query the GoLang API to retrieve a user's SMS history.

---

## Data Models (Structs)

The system relies on a unified data contract between the Java Producer and the Go Consumer. 

### `SMSRequest`
Represents the incoming payload for sending a message.
```go
type SMSRequest struct {
    UserID      string `json:"userId"`
    PhoneNumber string `json:"phoneNumber"`
    Message     string `json:"message"`
}
```

### `SMSResponse`
```go
type SMSResponse struct {
    TransactionID string `json:"transactionId,omitempty"`
    UserID        string `json:"userId"`
    PhoneNumber   string `json:"phoneNumber"`
    Message       string `json:"message"`
    Status        string `json:"status"` // e.g., "SUCCESS", "FAIL", "BLOCKED"
    Timestamp     string `json:"timestamp"`
}
```

## APIs
```go

// for sending sms(Java)

URL: /v1/sms/send
Method: POST
Content-Type: application/json


// for fetching sms history (Go)

URL: /v1/user/{User_id}/messages
Method: GET
```

## Instructions to Run it Locally
1. clone Both the Repos
2. Install all the dependencies for both
3. Run Redis-Server Locally
4. Start Docker Desktop and run Kafka through it
5. Configure the .env for Go repo
6. Run the Sprinboot app
7. Run the Go app (go run main.go)
8. Success , Now you can test those APIs

## Author
Niku Raj :)












