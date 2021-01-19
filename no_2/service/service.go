package service

import (
	"context"

	omdb "github.com/mghifariyusuf/stockbit_test.git/no_2/domain/omdb/repository"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
)

// Service ...
type Service interface {
	Search(ctx context.Context, r *SearchRequest) (e []entity.Movie, err error)
}

type service struct {
	omdbRepo omdb.Repository
}

// New ...
func New(
	omdbRepo omdb.Repository,
) Service {
	return &service{
		omdbRepo: omdbRepo,
	}
}
