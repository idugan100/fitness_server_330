package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin  bool   `json:"isAdmin"`
}
