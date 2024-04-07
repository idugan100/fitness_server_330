package models

type Activity struct {
	name string `json:"name" binding:"required"`
	id   int    `json:"id"`
}
