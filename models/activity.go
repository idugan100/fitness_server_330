package models

import "time"

type Activity struct {
	Name      string    `json:"name" binding:"required"`
	Id        int       `json:"id"`
	UserID    int       `json:"userID"`
	Intensity string    `json:"intensity" binding:"required"`
	Duration  int       `json:"duration" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
}
