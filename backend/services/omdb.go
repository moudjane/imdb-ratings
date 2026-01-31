package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/moudjane/imdb-ratings/backend/models"
)

func CalculateAverage(episodes []models.OmdbEpisode) string {
	var sum float64
	var count int
	for _, ep := range episodes {
		if val, err := strconv.ParseFloat(ep.ImdbRating, 64); err == nil {
			sum += val
			count++
		}
	}
	if count == 0 {
		return "N/A"
	}
	return fmt.Sprintf("%.1f", sum/float64(count))
}

func FetchAllSeasonsData(title string) ([]map[string]interface{}, string, error) {
	apiKey := os.Getenv("OMDB_API_KEY")
	var seasonsData []map[string]interface{}
	var allEpisodes []models.OmdbEpisode
	seasonNumber := 1

	for {
		safeTitle := url.QueryEscape(title)
		apiURL := fmt.Sprintf("http://www.omdbapi.com/?t=%s&Season=%d&apikey=%s", safeTitle, seasonNumber, apiKey)

		resp, err := http.Get(apiURL)
		if err != nil {
			return nil, "", err
		}
		defer resp.Body.Close()

		var result models.OmdbSeasonResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, "", err
		}
		if result.Response == "False" {
			break
		}

		seasonAvg := CalculateAverage(result.Episodes)
		allEpisodes = append(allEpisodes, result.Episodes...)

		seasonsData = append(seasonsData, map[string]interface{}{
			"Season":   result.Season,
			"Episodes": result.Episodes,
			"Average":  seasonAvg,
		})
		seasonNumber++
		if seasonNumber > 50 {
			break
		}
	}

	if len(seasonsData) == 0 {
		return nil, "", fmt.Errorf("série non trouvée")
	}

	globalAvg := CalculateAverage(allEpisodes)
	return seasonsData, globalAvg, nil
}
