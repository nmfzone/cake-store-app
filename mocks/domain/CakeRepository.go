// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/nmfzone/privy-cake-store/domain"
	mock "github.com/stretchr/testify/mock"
)

// CakeRepository is an autogenerated mock type for the CakeRepository type
type CakeRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: ctx, cursor, limit
func (_m *CakeRepository) FindAll(ctx context.Context, cursor string, limit int) ([]domain.Cake, string, error) {
	ret := _m.Called(ctx, cursor, limit)

	var r0 []domain.Cake
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []domain.Cake); ok {
		r0 = rf(ctx, cursor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Cake)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int) string); ok {
		r1 = rf(ctx, cursor, limit)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int) error); ok {
		r2 = rf(ctx, cursor, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: ctx, id
func (_m *CakeRepository) FindById(ctx context.Context, id uint64) (domain.Cake, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Cake
	if rf, ok := ret.Get(0).(func(context.Context, uint64) domain.Cake); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Cake)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByTitle provides a mock function with given fields: ctx, title
func (_m *CakeRepository) FindByTitle(ctx context.Context, title string) (domain.Cake, error) {
	ret := _m.Called(ctx, title)

	var r0 domain.Cake
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Cake); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(domain.Cake)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, cake
func (_m *CakeRepository) Remove(ctx context.Context, cake *domain.Cake) error {
	ret := _m.Called(ctx, cake)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Cake) error); ok {
		r0 = rf(ctx, cake)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, cake
func (_m *CakeRepository) Save(ctx context.Context, cake *domain.Cake) error {
	ret := _m.Called(ctx, cake)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Cake) error); ok {
		r0 = rf(ctx, cake)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCakeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCakeRepository creates a new instance of CakeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCakeRepository(t mockConstructorTestingTNewCakeRepository) *CakeRepository {
	mock := &CakeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
