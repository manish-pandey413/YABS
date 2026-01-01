package main

import (
	"context"
	"log"

	"github.com/manish-pandey413/YABS/internal/config"
	"github.com/manish-pandey413/YABS/internal/repository"
	"github.com/manish-pandey413/YABS/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	svr, err := server.New(cfg)
	if err != nil {
		log.Fatal("failed to initialize server")
	}

	_ = repository.NewRepositories(svr)
}
