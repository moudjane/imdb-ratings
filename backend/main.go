package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moudjane/imdb-ratings/backend/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/series/:title/season/:num", handlers.GetSeriesSeason)

	r.Run(":8080")
}
