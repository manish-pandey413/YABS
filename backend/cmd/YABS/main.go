package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/manish-pandey413/YABS/internal/config"
	"github.com/manish-pandey413/YABS/internal/handler"
	"github.com/manish-pandey413/YABS/internal/repository"
	"github.com/manish-pandey413/YABS/internal/router"
	"github.com/manish-pandey413/YABS/internal/server"
	"github.com/manish-pandey413/YABS/internal/service"
)

const DefaultContextTimeout = 30

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	svr, err := server.New(cfg)
	if err != nil {
		log.Println(fmt.Errorf("failed to initialize server, %w", err))
	}

	repos := repository.NewRepositories(svr)

	services, serviceError := service.NewServices(svr, repos)

	if serviceError != nil {
		fmt.Println(fmt.Errorf("Could not create services"))
	}

	handlers := handler.NewHandlers(svr, services)

	r := router.NewRouter(svr, handlers)
	svr.SetupHTTPServer(r)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	// Start server
	go func() {
		if err = svr.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeout*time.Second)

	if err = svr.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown")
	}
	stop()
	cancel()

	fmt.Println("server exited properly")
}
