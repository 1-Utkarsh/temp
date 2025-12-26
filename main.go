package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/1-Utkarsh/temp/api"
	"github.com/1-Utkarsh/temp/conf"
	db "github.com/1-Utkarsh/temp/store"
)

func main() {
	log.Default().Println("Application started")

	// initialize configuration
	conf.New()

	// connect to the database
	db.DbConnect()
	fmt.Println("Configuration and Database setup complete")

	// initialize api router
	r := router.InitRoutes()

	log.Default().Println("Listening on Port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Default().Fatal("Failed to start server:", err)
	}
}
