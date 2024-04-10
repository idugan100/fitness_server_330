package models

import "time"

type Notification struct {
	Id          int       `json:"id"`
	RecipientId int       `json:"recipientID"`
	Message     string    `json:"message"`
	Read        bool      `json:"read"`
	CreatedAt   time.Time `json:"createdAt"`
}
