// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	model "ad-server-project/src/domain/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// GetById provides a mock function with given fields: c, userId
func (_m *UserRepository) GetById(c context.Context, userId int) (model.User, error) {
	ret := _m.Called(c, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (model.User, error)); ok {
		return rf(c, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) model.User); ok {
		r0 = rf(c, userId)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(c, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReward provides a mock function with given fields: c, user
func (_m *UserRepository) UpdateReward(c context.Context, user model.User) error {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReward")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
