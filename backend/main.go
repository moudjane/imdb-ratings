package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moudjane/imdb-ratings/backend/handlers"
)

func main() {
	_ = godotenv.Load()
	r := gin.Default()

	r.LoadHTMLGlob("frontend/templates/*")
	r.Static("/static", "frontend/static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/series/:title", handlers.GetAllSeriesData)

	r.Run(":8080")
}
