package model

type Post struct {
	ID       int      `json:"id"`
	Owner_ID int      `json:"owner_id" validator:"required"`
	Content  string   `json:"content" validator:"required"`
	Comments []string `json:"comments"`
}
