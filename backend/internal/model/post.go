package model

type Post struct {
	ID             int      `json:"id"`
	Owner_Username string   `json:"owner_username" validator:"required"`
	Content        string   `json:"content" validator:"required"`
	Comments       []string `json:"comments"`
}
