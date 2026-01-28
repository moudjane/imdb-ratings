package handlers

import (
	"net/http"

	"github.com/moudjane/imdb-ratings/backend/services"

	"github.com/gin-gonic/gin"
)

func GetSeriesSeason(c *gin.Context) {
	title := c.Param("title")
	season := c.Param("num")

	data, err := services.FetchSeasonData(title, season)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur service"})
		return
	}

	if data.Response == "False" {
		c.JSON(http.StatusNotFound, gin.H{"error": data.Error})
		return
	}

	c.JSON(http.StatusOK, data)
}
