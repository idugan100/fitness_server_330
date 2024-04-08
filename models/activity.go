package models

import "time"

type Activity struct {
	Name      string `json:"name" binding:"required"`
	Id        int    `json:"id"`
	UserID    int
	Intensity string
	Duration  int
	Date      time.Time
}
