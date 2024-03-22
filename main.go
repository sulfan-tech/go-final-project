package main

import (
	"go-final-project/internal/delivery/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := ".env"
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	g := gin.Default()

	routes.RegisterRouter(g)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := ":" + port
	log.Printf("Server is running on port %s\n", port)
	if err := g.Run(address); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
