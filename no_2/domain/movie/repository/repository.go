package repository

import (
	"context"

	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
)

// Repository ...
type Repository interface {
	Upsert(ctx context.Context, e entity.Movie) (err error)
}
