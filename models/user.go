package models

type User struct {
	id       int    `json:"id"`
	userName string `json:"username" binding:"required"`
	password string `json:"password" binding:"required"`
	isAdmin  bool   `json:"isAdmin"`
}
