package main

import (
	"log"
	"net/http"
	"sayban/internal/config"
	"sayban/internal/router"
	"sayban/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.Load()

	userService := service.NewUserService()
	r := router.NewRouter(cfg, userService)

	log.Println("Server started on", cfg.ServerAddr)
	err := http.ListenAndServe(cfg.ServerAddr, r)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
