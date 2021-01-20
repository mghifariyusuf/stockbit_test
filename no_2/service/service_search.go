package service

import (
	"context"
	"log"
	"sync"

	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
)

// SearchRequest ...
type SearchRequest struct {
	SearchWord string
	Page       int
}

// Search ...
func (s *service) Search(ctx context.Context, r *SearchRequest) (e []entity.Movie, err error) {
	// call repo search
	results, err := s.movieRepo.Search(ctx, r.SearchWord, r.Page)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// get single detail movie in goroutine
	var wg sync.WaitGroup
	e = make([]entity.Movie, len(results))
	for i, v := range results {
		wg.Add(1)
		go func(i int, v entity.Movie) {
			defer wg.Done()
			detail, err := s.movieRepo.GetDetail(ctx, v.ImdbID)
			if err != nil {
				log.Println(err)
				return
			}
			e[i] = detail
		}(i, v)
	}
	wg.Wait()

	return e, nil
}
