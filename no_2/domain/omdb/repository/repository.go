package repository

import (
	"context"

	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
)

// Repository ...
type Repository interface {
	Search(ctx context.Context, searchWord string, page int) (e []entity.Movie, err error)
	GetDetail(ctx context.Context, id string) (e *entity.Movie, err error)
}
