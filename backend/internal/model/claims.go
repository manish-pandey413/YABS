package model

import "github.com/golang-jwt/jwt"

type Claims struct {
	User_id    int    `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email_id"`
	Authorized bool   `json:"authorized"`
	Expiration int64  `json:"exp"`
	jwt.StandardClaims
}
