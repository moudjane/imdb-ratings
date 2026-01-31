package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moudjane/imdb-ratings/backend/services"
)

func GetAllSeriesData(c *gin.Context) {
	title := c.Param("title")

	data, globalAvg, err := services.FetchAllSeasonsData(title)

	if err != nil {
		c.String(http.StatusNotFound, "Erreur : "+err.Error())
		return
	}

	isHX := c.GetHeader("HX-Request") == "true"

	if isHX {
		c.HTML(http.StatusOK, "all_seasons.html", gin.H{
			"Title":         title,
			"Seasons":       data,
			"GlobalAverage": globalAvg,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"InitialTitle": title,
		})
	}
}
