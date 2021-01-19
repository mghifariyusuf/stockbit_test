package handler

import (
	"context"
	"log"

	"github.com/mghifariyusuf/stockbit_test.git/no_2/delivery/grpc/schema"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/service"
)

// Handler ...
type Handler struct {
	service service.Service
}

// New ...
func New(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Search ...
func (h *Handler) Search(ctx context.Context, r *schema.SearchRequest) (*schema.SearchResponse, error) {
	// call service for search
	e, err := h.service.Search(ctx, &service.SearchRequest{
		SearchWord: r.SearchWord,
		Page:       int(r.Pagination),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// convert from entity to grpc schema
	mv := make([]*schema.Movie, 0, len(e))
	for _, v := range e {
		ratings := make([]*schema.Rating, 0, len(v.Ratings))
		for _, v := range v.Ratings {
			ratings = append(ratings, &schema.Rating{
				Source: v.Source,
				Value:  v.Value,
			})
		}

		mv = append(mv, &schema.Movie{
			Title:      v.Title,
			Year:       v.Year,
			Rated:      v.Rated,
			Released:   v.Released,
			Runtime:    v.Runtime,
			Genre:      v.Genre,
			Director:   v.Director,
			Writer:     v.Writer,
			Actors:     v.Actors,
			Plot:       v.Plot,
			Language:   v.Language,
			Country:    v.Country,
			Awards:     v.Awards,
			Poster:     v.Poster,
			Ratings:    ratings,
			Metascore:  v.Metascore,
			ImdbRating: v.ImdbRating,
			ImdbVotes:  v.ImdbVotes,
			ImdbID:     v.ImdbID,
			Type:       v.Type,
			Dvd:        v.DVD,
			BoxOffice:  v.BoxOffice,
			Production: v.Production,
			Website:    v.Website,
		})
	}

	return &schema.SearchResponse{Search: mv}, nil
}
