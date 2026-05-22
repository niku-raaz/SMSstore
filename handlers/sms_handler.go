package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"SMSstore/models"
	"SMSstore/service"
)

func TestInsert(w http.ResponseWriter, r *http.Request) {

	sms := models.SMS{
		UserID:    1,
		Recipient: "9999999999",
		Message:   "Hello Mongo",
		Status:    "SUCCESS",
		CreatedAt: time.Now(),
	}

	err := service.SaveSMS(sms)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Inserted Successfully"))
}

func GetMessagesByUser(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.PathValue("userId")
	fmt.Println("userId:", userIDStr)

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	fmt.Println("parsed userId:", userID)

	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	messages, err := service.GetMessagesByUserID(userID)

	for _, mssg := range messages {
		fmt.Println(mssg)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(messages)
}
