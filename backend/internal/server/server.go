package server

import (
	"github.com/manish-pandey413/YABS/internal/config"
	"github.com/manish-pandey413/YABS/internal/database"
)

type Server struct {
	Config *config.Config
	DB     *database.Database
}

func New(cfg *config.Config) (*Server, error) {
	db, err := database.New(cfg)
	if err != nil {
		return nil, err
	}
	server := &Server{
		Config: cfg,
		DB:     db,
	}
	return server, nil
}
