// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entities "gorepair-rest-api/src/users/entities"

	mock "github.com/stretchr/testify/mock"
)

// UserMysqlRepositoryInterface is an autogenerated mock type for the UserMysqlRepositoryInterface type
type UserMysqlRepositoryInterface struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: email
func (_m *UserMysqlRepositoryInterface) FindByEmail(email string) *entities.Users {
	ret := _m.Called(email)

	var r0 *entities.Users
	if rf, ok := ret.Get(0).(func(string) *entities.Users); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Users)
		}
	}

	return r0
}

// GetAddress provides a mock function with given fields: id
func (_m *UserMysqlRepositoryInterface) GetAddress(id uint64) (*entities.UserAddress, error) {
	ret := _m.Called(id)

	var r0 *entities.UserAddress
	if rf, ok := ret.Get(0).(func(uint64) *entities.UserAddress); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserAddress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: username
func (_m *UserMysqlRepositoryInterface) GetUser(username string) (*entities.Users, error) {
	ret := _m.Called(username)

	var r0 *entities.Users
	if rf, ok := ret.Get(0).(func(string) *entities.Users); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: payload, street
func (_m *UserMysqlRepositoryInterface) Register(payload *entities.Users, street string) (*entities.Users, error) {
	ret := _m.Called(payload, street)

	var r0 *entities.Users
	if rf, ok := ret.Get(0).(func(*entities.Users, string) *entities.Users); ok {
		r0 = rf(payload, street)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Users, string) error); ok {
		r1 = rf(payload, street)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccount provides a mock function with given fields: payload, id
func (_m *UserMysqlRepositoryInterface) UpdateAccount(payload *entities.Users, id uint64) (*entities.Users, error) {
	ret := _m.Called(payload, id)

	var r0 *entities.Users
	if rf, ok := ret.Get(0).(func(*entities.Users, uint64) *entities.Users); ok {
		r0 = rf(payload, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Users, uint64) error); ok {
		r1 = rf(payload, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAddress provides a mock function with given fields: payload, id
func (_m *UserMysqlRepositoryInterface) UpdateAddress(payload *entities.UserAddress, id uint64) (*entities.UserAddress, error) {
	ret := _m.Called(payload, id)

	var r0 *entities.UserAddress
	if rf, ok := ret.Get(0).(func(*entities.UserAddress, uint64) *entities.UserAddress); ok {
		r0 = rf(payload, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserAddress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.UserAddress, uint64) error); ok {
		r1 = rf(payload, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}