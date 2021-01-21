package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
	mockrepo "github.com/mghifariyusuf/stockbit_test.git/no_2/mocks/repository"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/testdata"
	"github.com/stretchr/testify/require"
)

func TestGetDetail(t *testing.T) {
	var (
		id          = "abcd"
		entityMovie entity.Movie
	)
	testdata.GoldenJSONUnmarshal(t, "movie", &entityMovie)

	tests := map[string]struct {
		movieDetail   testdata.FuncCaller
		responseMovie entity.Movie
		responseError error
	}{
		"success": {
			movieDetail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{context.TODO(), id},
				Output:   []interface{}{entityMovie, nil},
			},
			responseMovie: entityMovie,
			responseError: nil,
		},
		"error-not-found": {
			movieDetail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{context.TODO(), id},
				Output:   []interface{}{entity.Movie{}, sql.ErrNoRows},
			},
			responseMovie: entity.Movie{},
			responseError: errors.New("sql: no rows in result set"),
		},
		"error": {
			movieDetail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{context.TODO(), id},
				Output:   []interface{}{entity.Movie{}, fmt.Errorf("error")},
			},
			responseMovie: entity.Movie{},
			responseError: errors.New("error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			omdbRepo := new(mockrepo.Repository)
			serviceMock := New(
				omdbRepo,
			)

			omdbRepo.On("GetDetail", test.movieDetail.Input...).
				Return(test.movieDetail.Output...).
				Once()

			responseMovie, err := serviceMock.GetDetail(context.TODO(), &GetDetailRequest{ID: id})
			if err != nil {
				require.Error(t, err)
				require.Equal(t, responseMovie.ImdbID, test.responseMovie.ImdbID)
				require.Equal(t, err, test.responseError)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, responseMovie)
			require.Equal(t, responseMovie.ImdbID, test.responseMovie.ImdbID)
		})
	}
}
