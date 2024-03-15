// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	mock "github.com/stretchr/testify/mock"
)

// EmployeeRepo is an autogenerated mock type for the EmployeeRepo type
type EmployeeRepo struct {
	mock.Mock
}

type EmployeeRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *EmployeeRepo) EXPECT() *EmployeeRepo_Expecter {
	return &EmployeeRepo_Expecter{mock: &_m.Mock}
}

// GetEmployeeAndBalance provides a mock function with given fields: employeeNumber, companyID
func (_m *EmployeeRepo) GetEmployeeAndBalance(employeeNumber string, companyID string) (entity.Employee, entity.Balance, error) {
	ret := _m.Called(employeeNumber, companyID)

	var r0 entity.Employee
	var r1 entity.Balance
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (entity.Employee, entity.Balance, error)); ok {
		return rf(employeeNumber, companyID)
	}
	if rf, ok := ret.Get(0).(func(string, string) entity.Employee); ok {
		r0 = rf(employeeNumber, companyID)
	} else {
		r0 = ret.Get(0).(entity.Employee)
	}

	if rf, ok := ret.Get(1).(func(string, string) entity.Balance); ok {
		r1 = rf(employeeNumber, companyID)
	} else {
		r1 = ret.Get(1).(entity.Balance)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(employeeNumber, companyID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EmployeeRepo_GetEmployeeAndBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEmployeeAndBalance'
type EmployeeRepo_GetEmployeeAndBalance_Call struct {
	*mock.Call
}

// GetEmployeeAndBalance is a helper method to define mock.On call
//   - employeeNumber string
//   - companyID string
func (_e *EmployeeRepo_Expecter) GetEmployeeAndBalance(employeeNumber interface{}, companyID interface{}) *EmployeeRepo_GetEmployeeAndBalance_Call {
	return &EmployeeRepo_GetEmployeeAndBalance_Call{Call: _e.mock.On("GetEmployeeAndBalance", employeeNumber, companyID)}
}

func (_c *EmployeeRepo_GetEmployeeAndBalance_Call) Run(run func(employeeNumber string, companyID string)) *EmployeeRepo_GetEmployeeAndBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *EmployeeRepo_GetEmployeeAndBalance_Call) Return(_a0 entity.Employee, _a1 entity.Balance, _a2 error) *EmployeeRepo_GetEmployeeAndBalance_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EmployeeRepo_GetEmployeeAndBalance_Call) RunAndReturn(run func(string, string) (entity.Employee, entity.Balance, error)) *EmployeeRepo_GetEmployeeAndBalance_Call {
	_c.Call.Return(run)
	return _c
}

// GetEmployeeNumberFromRFID provides a mock function with given fields: rfid, companyID
func (_m *EmployeeRepo) GetEmployeeNumberFromRFID(rfid string, companyID string) (string, error) {
	ret := _m.Called(rfid, companyID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(rfid, companyID)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(rfid, companyID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(rfid, companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EmployeeRepo_GetEmployeeNumberFromRFID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEmployeeNumberFromRFID'
type EmployeeRepo_GetEmployeeNumberFromRFID_Call struct {
	*mock.Call
}

// GetEmployeeNumberFromRFID is a helper method to define mock.On call
//   - rfid string
//   - companyID string
func (_e *EmployeeRepo_Expecter) GetEmployeeNumberFromRFID(rfid interface{}, companyID interface{}) *EmployeeRepo_GetEmployeeNumberFromRFID_Call {
	return &EmployeeRepo_GetEmployeeNumberFromRFID_Call{Call: _e.mock.On("GetEmployeeNumberFromRFID", rfid, companyID)}
}

func (_c *EmployeeRepo_GetEmployeeNumberFromRFID_Call) Run(run func(rfid string, companyID string)) *EmployeeRepo_GetEmployeeNumberFromRFID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *EmployeeRepo_GetEmployeeNumberFromRFID_Call) Return(_a0 string, _a1 error) *EmployeeRepo_GetEmployeeNumberFromRFID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EmployeeRepo_GetEmployeeNumberFromRFID_Call) RunAndReturn(run func(string, string) (string, error)) *EmployeeRepo_GetEmployeeNumberFromRFID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewEmployeeRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmployeeRepo creates a new instance of EmployeeRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmployeeRepo(t mockConstructorTestingTNewEmployeeRepo) *EmployeeRepo {
	mock := &EmployeeRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
