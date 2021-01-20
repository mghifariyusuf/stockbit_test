package movie

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

	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
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

type movie struct {
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
	Ratings    []rating `json:"Ratings"`
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

type rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

var result struct {
	Search       []movie `json:"Search"`
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
	// create new request
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", h.cfg.BaseURL, h.cfg.Key, searchWord, page)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// do the request
	response, err := h.http.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	// read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// unmarshal to struct
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// if status not 200 or return response false
	if response.StatusCode != 200 || result.Response == "False" {
		log.Println(result.Error)
		io.Copy(ioutil.Discard, response.Body)
		return nil, errors.New(result.Error)
	}

	// collecting imdb id
	for _, v := range result.Search {
		e = append(e, entity.Movie{
			ImdbID: v.ImdbID,
		})
	}

	return e, nil
}

// GetDetail ...
func (h *HTTP) GetDetail(ctx context.Context, id string) (e entity.Movie, err error) {
	// new request
	url := fmt.Sprintf("%s/?apikey=%s&i=%s", h.cfg.BaseURL, h.cfg.Key, id)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// do request
	response, err := h.http.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	// read body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// unmarshal to struct
	var mv movie
	err = json.Unmarshal(body, &mv)
	if err != nil {
		log.Println(err)
		return
	}

	// if status not 200 or return response false
	if response.StatusCode != 200 || result.Response == "False" {
		log.Println(result.Error)
		io.Copy(ioutil.Discard, response.Body)
		return
	}

	// convert result to entity
	ratings := make([]entity.Rating, 0, len(mv.Ratings))
	for _, v := range mv.Ratings {
		rating := entity.Rating{
			Source: v.Source,
			Value:  v.Value,
		}
		ratings = append(ratings, rating)
	}

	e = entity.Movie{
		Title:      mv.Title,
		Year:       mv.Year,
		Rated:      mv.Rated,
		Released:   mv.Released,
		Runtime:    mv.Runtime,
		Genre:      mv.Genre,
		Director:   mv.Director,
		Writer:     mv.Writer,
		Actors:     mv.Actors,
		Plot:       mv.Plot,
		Language:   mv.Language,
		Country:    mv.Country,
		Awards:     mv.Awards,
		Poster:     mv.Poster,
		Ratings:    ratings,
		Metascore:  mv.Metascore,
		ImdbRating: mv.ImdbRating,
		ImdbVotes:  mv.ImdbVotes,
		ImdbID:     mv.ImdbID,
		Type:       mv.Type,
		DVD:        mv.DVD,
		BoxOffice:  mv.BoxOffice,
		Production: mv.Production,
		Website:    mv.Website,
	}

	return e, nil
}
