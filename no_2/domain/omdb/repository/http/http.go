package omdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HTTP ...
type HTTP struct {
	http *http.Client
	cfg  *Config
}

// Config ...
type Config struct {
	BaseURL string
	Key     string
}

// Movie ...
type Movie struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
}

// Rating ...
type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

var result struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
	Error        string  `json:"Error"`
}

// New ...
func New(cfg *Config) *HTTP {
	return &HTTP{
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
		cfg: cfg,
	}
}

// Search ...
func (h *HTTP) Search(ctx context.Context, searchWord string, page int) (e []entity.Movie, err error) {
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", h.cfg.BaseURL, h.cfg.Key, searchWord, page)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response, err := h.http.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if response.StatusCode != 200 || result.Response == "False" {
		io.Copy(ioutil.Discard, response.Body)
		return nil, errors.New(result.Error)
	}

	return e, nil
}
