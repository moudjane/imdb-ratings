package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/moudjane/imdb-ratings/backend/models"
)

func FetchSeasonData(title, season string) (*models.OmdbSeasonResponse, error) {
	apiKey := os.Getenv("OMDB_API_KEY")

	safeTitle := url.QueryEscape(title)
	apiURL := fmt.Sprintf("http://www.omdbapi.com/?t=%s&Season=%s&apikey=%s", safeTitle, season, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.OmdbSeasonResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
