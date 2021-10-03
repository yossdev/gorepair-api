// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entities "gorepair-rest-api/src/workshops/entities"

	mock "github.com/stretchr/testify/mock"
)

// WorkshopMysqlRepositoryInterface is an autogenerated mock type for the WorkshopMysqlRepositoryInterface type
type WorkshopMysqlRepositoryInterface struct {
	mock.Mock
}

// DeleteServices provides a mock function with given fields: id, servicesId
func (_m *WorkshopMysqlRepositoryInterface) DeleteServices(id uint64, servicesId uint64) error {
	ret := _m.Called(id, servicesId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) error); ok {
		r0 = rf(id, servicesId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByEmail provides a mock function with given fields: email
func (_m *WorkshopMysqlRepositoryInterface) FindByEmail(email string) *entities.Workshops {
	ret := _m.Called(email)

	var r0 *entities.Workshops
	if rf, ok := ret.Get(0).(func(string) *entities.Workshops); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Workshops)
		}
	}

	return r0
}

// GetAddress provides a mock function with given fields: id
func (_m *WorkshopMysqlRepositoryInterface) GetAddress(id uint64) (*entities.WorkshopAddress, error) {
	ret := _m.Called(id)

	var r0 *entities.WorkshopAddress
	if rf, ok := ret.Get(0).(func(uint64) *entities.WorkshopAddress); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.WorkshopAddress)
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

// GetWorkshop provides a mock function with given fields: username
func (_m *WorkshopMysqlRepositoryInterface) GetWorkshop(username string) (*entities.Workshops, error) {
	ret := _m.Called(username)

	var r0 *entities.Workshops
	if rf, ok := ret.Get(0).(func(string) *entities.Workshops); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Workshops)
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

// Register provides a mock function with given fields: payload, street, description
func (_m *WorkshopMysqlRepositoryInterface) Register(payload *entities.Workshops, street string, description string) (*entities.Workshops, error) {
	ret := _m.Called(payload, street, description)

	var r0 *entities.Workshops
	if rf, ok := ret.Get(0).(func(*entities.Workshops, string, string) *entities.Workshops); ok {
		r0 = rf(payload, street, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Workshops)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Workshops, string, string) error); ok {
		r1 = rf(payload, street, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServicesNew provides a mock function with given fields: payload, id
func (_m *WorkshopMysqlRepositoryInterface) ServicesNew(payload *entities.Services, id uint64) (*entities.Services, error) {
	ret := _m.Called(payload, id)

	var r0 *entities.Services
	if rf, ok := ret.Get(0).(func(*entities.Services, uint64) *entities.Services); ok {
		r0 = rf(payload, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Services)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Services, uint64) error); ok {
		r1 = rf(payload, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccount provides a mock function with given fields: payload, id
func (_m *WorkshopMysqlRepositoryInterface) UpdateAccount(payload *entities.Workshops, id uint64) (*entities.Workshops, error) {
	ret := _m.Called(payload, id)

	var r0 *entities.Workshops
	if rf, ok := ret.Get(0).(func(*entities.Workshops, uint64) *entities.Workshops); ok {
		r0 = rf(payload, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Workshops)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Workshops, uint64) error); ok {
		r1 = rf(payload, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAddress provides a mock function with given fields: payload, id
func (_m *WorkshopMysqlRepositoryInterface) UpdateAddress(payload *entities.WorkshopAddress, id uint64) (*entities.WorkshopAddress, error) {
	ret := _m.Called(payload, id)

	var r0 *entities.WorkshopAddress
	if rf, ok := ret.Get(0).(func(*entities.WorkshopAddress, uint64) *entities.WorkshopAddress); ok {
		r0 = rf(payload, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.WorkshopAddress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.WorkshopAddress, uint64) error); ok {
		r1 = rf(payload, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDescription provides a mock function with given fields: payload, id
func (_m *WorkshopMysqlRepositoryInterface) UpdateDescription(payload *entities.Descriptions, id uint64) (*entities.Descriptions, error) {
	ret := _m.Called(payload, id)

	var r0 *entities.Descriptions
	if rf, ok := ret.Get(0).(func(*entities.Descriptions, uint64) *entities.Descriptions); ok {
		r0 = rf(payload, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Descriptions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Descriptions, uint64) error); ok {
		r1 = rf(payload, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateServices provides a mock function with given fields: payload, id, servicesId
func (_m *WorkshopMysqlRepositoryInterface) UpdateServices(payload *entities.Services, id uint64, servicesId uint64) (*entities.Services, error) {
	ret := _m.Called(payload, id, servicesId)

	var r0 *entities.Services
	if rf, ok := ret.Get(0).(func(*entities.Services, uint64, uint64) *entities.Services); ok {
		r0 = rf(payload, id, servicesId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Services)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entities.Services, uint64, uint64) error); ok {
		r1 = rf(payload, id, servicesId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}