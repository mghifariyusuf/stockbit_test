// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetDetail provides a mock function with given fields: ctx, id
func (_m *Repository) GetDetail(ctx context.Context, id string) (entity.Movie, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Movie
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Movie); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Movie)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, searchWord, page
func (_m *Repository) Search(ctx context.Context, searchWord string, page int) ([]entity.Movie, error) {
	ret := _m.Called(ctx, searchWord, page)

	var r0 []entity.Movie
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []entity.Movie); ok {
		r0 = rf(ctx, searchWord, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, searchWord, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
