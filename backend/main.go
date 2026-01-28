package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moudjane/imdb-ratings/backend/handlers"
)

func main() {
	_ = godotenv.Load()

	r := gin.Default()

	r.GET("/series/:title/season/:num", handlers.GetSeriesSeason)

	r.Run(":8080")
}
