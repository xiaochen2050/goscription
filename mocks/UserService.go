// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/kecci/goscription/models"
	mock "github.com/stretchr/testify/mock"

	service "github.com/kecci/goscription/internal/service"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UserService) GetByEmail(ctx context.Context, email string) (models.User, error) {
	ret := _m.Called(ctx, email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserService) GetByID(ctx context.Context, id int64) (models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *UserService) Store(_a0 context.Context, _a1 service.UserParam) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, service.UserParam) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
