// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	model "ad-server-project/src/domain/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AdvertisementUsecase is an autogenerated mock type for the AdvertisementUsecase type
type AdvertisementUsecase struct {
	mock.Mock
}

// GetByCountryAndGender provides a mock function with given fields: c, userId, userGender, userCountry
func (_m *AdvertisementUsecase) GetByCountryAndGender(c context.Context, userId int, userGender string, userCountry string) ([]model.AdvertisementMinInfo, error) {
	ret := _m.Called(c, userId, userGender, userCountry)

	if len(ret) == 0 {
		panic("no return value specified for GetByCountryAndGender")
	}

	var r0 []model.AdvertisementMinInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) ([]model.AdvertisementMinInfo, error)); ok {
		return rf(c, userId, userGender, userCountry)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) []model.AdvertisementMinInfo); ok {
		r0 = rf(c, userId, userGender, userCountry)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AdvertisementMinInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, string) error); ok {
		r1 = rf(c, userId, userGender, userCountry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReward provides a mock function with given fields: c, id, reward
func (_m *AdvertisementUsecase) UpdateReward(c context.Context, id int, reward int) error {
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

// NewAdvertisementUsecase creates a new instance of AdvertisementUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdvertisementUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdvertisementUsecase {
	mock := &AdvertisementUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
