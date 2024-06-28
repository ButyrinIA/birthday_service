package main

import (
	"log"
	"net/http"

	"rutube/config"
	"rutube/internal/database"
	"rutube/internal/handlers"
	"rutube/internal/repository"
	"rutube/internal/routes"
	"rutube/internal/service"
)

func main() {
	config.Init()
	database.InitDB()

	userRepo := repository.NewUserRepository(database.DB)
	birthdayService := service.NewBirthdayService(userRepo)
	authService := service.NewAuthService(userRepo)
	birthdayHandler := handlers.NewBirthdayHandler(birthdayService, authService)

	router := routes.RegisterRoutes(birthdayHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
