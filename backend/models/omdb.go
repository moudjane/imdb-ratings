package models

type OmdbSeasonResponse struct {
	Title    string        `json:"Title"`
	Season   string        `json:"Season"`
	Episodes []OmdbEpisode `json:"Episodes"`
	Response string        `json:"Response"`
	Error    string        `json:"Error"`
}

type OmdbEpisode struct {
	Title      string `json:"Title"`
	Episode    string `json:"Episode"`
	ImdbRating string `json:"imdbRating"`
}
