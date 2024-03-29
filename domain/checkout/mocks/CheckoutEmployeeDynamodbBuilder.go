// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	entity "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	mock "github.com/stretchr/testify/mock"
)

// CheckoutEmployeeDynamodbBuilder is an autogenerated mock type for the CheckoutEmployeeDynamodbBuilder type
type CheckoutEmployeeDynamodbBuilder struct {
	mock.Mock
}

type CheckoutEmployeeDynamodbBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *CheckoutEmployeeDynamodbBuilder) EXPECT() *CheckoutEmployeeDynamodbBuilder_Expecter {
	return &CheckoutEmployeeDynamodbBuilder_Expecter{mock: &_m.Mock}
}

// BuildUpdateBalanceRequest provides a mock function with given fields: tableName, bal, newBal
func (_m *CheckoutEmployeeDynamodbBuilder) BuildUpdateBalanceRequest(tableName string, bal entity.Balance, newBal float64) *dynamodb.Update {
	ret := _m.Called(tableName, bal, newBal)

	var r0 *dynamodb.Update
	if rf, ok := ret.Get(0).(func(string, entity.Balance, float64) *dynamodb.Update); ok {
		r0 = rf(tableName, bal, newBal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.Update)
		}
	}

	return r0
}

// CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BuildUpdateBalanceRequest'
type CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call struct {
	*mock.Call
}

// BuildUpdateBalanceRequest is a helper method to define mock.On call
//   - tableName string
//   - bal entity.Balance
//   - newBal float64
func (_e *CheckoutEmployeeDynamodbBuilder_Expecter) BuildUpdateBalanceRequest(tableName interface{}, bal interface{}, newBal interface{}) *CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call {
	return &CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call{Call: _e.mock.On("BuildUpdateBalanceRequest", tableName, bal, newBal)}
}

func (_c *CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call) Run(run func(tableName string, bal entity.Balance, newBal float64)) *CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(entity.Balance), args[2].(float64))
	})
	return _c
}

func (_c *CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call) Return(_a0 *dynamodb.Update) *CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call) RunAndReturn(run func(string, entity.Balance, float64) *dynamodb.Update) *CheckoutEmployeeDynamodbBuilder_BuildUpdateBalanceRequest_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCheckoutEmployeeDynamodbBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewCheckoutEmployeeDynamodbBuilder creates a new instance of CheckoutEmployeeDynamodbBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCheckoutEmployeeDynamodbBuilder(t mockConstructorTestingTNewCheckoutEmployeeDynamodbBuilder) *CheckoutEmployeeDynamodbBuilder {
	mock := &CheckoutEmployeeDynamodbBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
