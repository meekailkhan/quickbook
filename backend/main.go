package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/meekailkhan/quick-book/internals/api"
)

const (
	PORT = ":3030"
)

func main() {
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		os.Create(".env")
	}
	if err := godotenv.Load(); err != nil {
		panic("could not load the env")
	}
	databaseUrl := os.Getenv("ENV_DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("database url not set in env")
	}

	fmt.Printf("datbase url is: %s\n", databaseUrl)

	router := api.SetupRouter()

	srv := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	go func() {
		log.Printf("server start on port%s\n", PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could start server %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGCHLD)
	val := <-quit
	fmt.Printf("shuting down server with system call: %d\n", val)

}
