package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validator:"required"`
	Email    string `json:"email" validator:"required"`
	Password string `json:"password" validator:"required"`
}
