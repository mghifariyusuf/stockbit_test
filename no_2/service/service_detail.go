package service

import (
	"context"
	"log"

	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
)

// GetDetailRequest ...
type GetDetailRequest struct {
	ID string
}

// GetDetail ...
func (s *service) GetDetail(ctx context.Context, r *GetDetailRequest) (e entity.Movie, err error) {
	// get data from omdb
	e, err = s.omdbRepo.GetDetail(ctx, r.ID)
	if err != nil {
		log.Println(err)
		return e, err
	}

	// insert db
	err = s.movieRepo.Upsert(ctx, e)
	if err != nil {
		log.Println(err)
		return
	}

	return e, nil
}
