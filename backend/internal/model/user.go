package model

type User struct {
	Username string `json:"username" validator:"required"`
	Email    string `json:"email" validator:"required"`
	Password string `json:"password" validator:"required"`
}
