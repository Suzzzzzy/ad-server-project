// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	model "ad-server-project/src/domain/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AdvertisementRepository is an autogenerated mock type for the AdvertisementRepository type
type AdvertisementRepository struct {
	mock.Mock
}

// GetByCountryAndGender provides a mock function with given fields: c, user
func (_m *AdvertisementRepository) GetByCountryAndGender(c context.Context, user *model.User) ([]model.Advertisement, error) {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for GetByCountryAndGender")
	}

	var r0 []model.Advertisement
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) ([]model.Advertisement, error)); ok {
		return rf(c, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) []model.Advertisement); ok {
		r0 = rf(c, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Advertisement)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(c, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: c, id
func (_m *AdvertisementRepository) GetById(c context.Context, id int) (model.Advertisement, error) {
	ret := _m.Called(c, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 model.Advertisement
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (model.Advertisement, error)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) model.Advertisement); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(model.Advertisement)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReward provides a mock function with given fields: c, id, reward
func (_m *AdvertisementRepository) UpdateReward(c context.Context, id int, reward int) error {
	ret := _m.Called(c, id, reward)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReward")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(c, id, reward)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdvertisementRepository creates a new instance of AdvertisementRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdvertisementRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdvertisementRepository {
	mock := &AdvertisementRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
