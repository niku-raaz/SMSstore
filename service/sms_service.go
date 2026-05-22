package service

import (
	"SMSstore/models"
	"SMSstore/repository"
)

func SaveSMS(sms models.SMS) error {
	return repository.SaveSMS(sms)
}

func GetMessagesByUserID(userID int64) ([]models.SMS, error) {
	return repository.GetMessagesByUserID(userID)
}
