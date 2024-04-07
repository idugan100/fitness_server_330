package models

import "time"

type Notification struct {
	Id          int       `json:"id"`
	RecipientId int       `json:"recipientID"`
	Message     string    `json:"message" binding:"required"`
	Read        bool      `json:"read" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" binding:"required"`
}
