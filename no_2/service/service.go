package service

import (
	"context"

	movie "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/movie/repository"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
)

// Service ...
type Service interface {
	Search(ctx context.Context, r *SearchRequest) (e []entity.Movie, err error)
	GetDetail(ctx context.Context, r *GetDetailRequest) (e entity.Movie, err error)
}

type service struct {
	movieRepo movie.Repository
}

// New ...
func New(
	movieRepo movie.Repository,
) Service {
	return &service{
		movieRepo: movieRepo,
	}
}
