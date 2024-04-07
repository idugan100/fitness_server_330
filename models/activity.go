package models

import "time"

type Activity struct {
	name      string `json:"name" binding:"required"`
	id        int    `json:"id"`
	userID    int
	intensity string
	duration  int
	date      time.Time
}
