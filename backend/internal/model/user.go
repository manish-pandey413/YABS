package model

type UserCred struct {
	Username string `json:"username" validator:"required"`
	Password string `json:"password" validator:"required"`
}

type User struct {
	User_id int `json:"user_id"`
	UserCred
	Email string `json:"email" validator:"required"`
}
