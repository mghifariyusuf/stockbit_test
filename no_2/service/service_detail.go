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
	// call repo
	e, err = s.movieRepo.GetDetail(ctx, r.ID)
	if err != nil {
		log.Println(err)
		return e, err
	}

	return e, nil
}
