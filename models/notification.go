package models

import "time"

type Notification struct {
	id          int       `json:"id"`
	recipientId int       `json:"recipientID"`
	message     string    `json:"message" binding:"required"`
	read        bool      `json:"read" binding:"required"`
	createdAt   time.Time `json:"createdAt" binding:"required"`
}
