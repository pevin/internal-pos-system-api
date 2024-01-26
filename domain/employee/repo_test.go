package employee_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pevin/internal-pos-service-api/domain/employee"
	entityEmployee "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockedDynamodbClient struct {
	mock.Mock
}

func (mr *mockedDynamodbClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	args := mr.Called(input)
	return args.Get(0).(*dynamodb.QueryOutput), args.Error(1)
}

func (mr *mockedDynamodbClient) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	args := mr.Called(input)
	return args.Get(0).(*dynamodb.BatchGetItemOutput), args.Error(1)
}

type mockEmployeeDynamodbBuilder struct {
	mock.Mock
}

func (b *mockEmployeeDynamodbBuilder) BuildGetRFIDQueryInput(employee entityEmployee.Employee, tableName string) (*dynamodb.QueryInput, error) {
	args := b.Called(employee, tableName)
	return args.Get(0).(*dynamodb.QueryInput), args.Error(1)
}

func TestGetEmployeeNumberFromRFID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Given a valid parameters and mock behaviors
		rfid := "123-123"
		companyID := "test-company-id"

		tableName := "test-table-name"
		dClient := new(mockedDynamodbClient)
		builder := new(mockEmployeeDynamodbBuilder)
		emp := entityEmployee.Employee{
			RFID:      rfid,
			CompanyID: companyID,
		}
		queryInput := &dynamodb.QueryInput{}
		builder.On("BuildGetRFIDQueryInput", emp, tableName).Return(queryInput, nil)

		empWithEmployeeNumber := entityEmployee.Employee{
			RFID:           rfid,
			CompanyID:      companyID,
			EmployeeNumber: "test-employee-number",
		}
		av, err := empWithEmployeeNumber.ToItem()
		require.NoError(t, err)
		queryOutput := &dynamodb.QueryOutput{
			Items: []map[string]*dynamodb.AttributeValue{
				av,
			},
		}
		dClient.On("Query", queryInput).Return(queryOutput, nil)

		repo := employee.NewRepo(employee.RepoOpt{
			TableName:               tableName,
			DynamodbClient:          dClient,
			EmployeeDynamodbBuilder: builder,
		})

		// When GetEmployeeNumberFromRFID is called
		actual, err := repo.GetEmployeeNumberFromRFID(rfid, companyID)

		// Then there should be no error and employee number is correct
		require.NoError(t, err)
		assert.Equal(t, "test-employee-number", actual)
	})
	t.Run("not found", func(t *testing.T) {
		// Given a dynamodb client that returns empty query output
		rfid := "123-123"
		companyID := "test-company-id"

		tableName := "test-table-name"
		dClient := new(mockedDynamodbClient)
		builder := new(mockEmployeeDynamodbBuilder)
		emp := entityEmployee.Employee{
			RFID:      rfid,
			CompanyID: companyID,
		}
		queryInput := &dynamodb.QueryInput{}
		builder.On("BuildGetRFIDQueryInput", emp, tableName).Return(queryInput, nil)

		queryOutput := &dynamodb.QueryOutput{
			Items: []map[string]*dynamodb.AttributeValue{},
		}
		dClient.On("Query", queryInput).Return(queryOutput, nil)

		repo := employee.NewRepo(employee.RepoOpt{
			TableName:               tableName,
			DynamodbClient:          dClient,
			EmployeeDynamodbBuilder: builder,
		})

		// When GetEmployeeNumberFromRFID is called
		actual, err := repo.GetEmployeeNumberFromRFID(rfid, companyID)

		// Then there should be no error and employee number is empty
		require.NoError(t, err)
		assert.Equal(t, "", actual)

	})
	t.Run("dynamodb client error", func(t *testing.T) {
		// Given a dynamodb client that returns empty query output
		rfid := "123-123"
		companyID := "test-company-id"

		tableName := "test-table-name"
		dClient := new(mockedDynamodbClient)
		builder := new(mockEmployeeDynamodbBuilder)
		emp := entityEmployee.Employee{
			RFID:      rfid,
			CompanyID: companyID,
		}
		queryInput := &dynamodb.QueryInput{}
		builder.On("BuildGetRFIDQueryInput", emp, tableName).Return(queryInput, nil)

		queryOutput := &dynamodb.QueryOutput{
			Items: []map[string]*dynamodb.AttributeValue{},
		}
		dClient.On("Query", queryInput).Return(queryOutput, errors.New("test error"))

		repo := employee.NewRepo(employee.RepoOpt{
			TableName:               tableName,
			DynamodbClient:          dClient,
			EmployeeDynamodbBuilder: builder,
		})

		// When GetEmployeeNumberFromRFID is called
		_, err := repo.GetEmployeeNumberFromRFID(rfid, companyID)

		// Then there should be an error
		assert.Error(t, err)
	})
}

func TestGetEmployeeAndBalance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Given a valid input and mock behavior
		employeeNumber := "test-employee-number"
		companyID := "test-company-id"

		tableName := "test-table-name"
		dClient := new(mockedDynamodbClient)
		builder := new(mockEmployeeDynamodbBuilder)

		repo := employee.NewRepo(employee.RepoOpt{
			TableName:               tableName,
			DynamodbClient:          dClient,
			EmployeeDynamodbBuilder: builder,
		})

		mockBal := entityEmployee.Balance{
			CompanyID:      companyID,
			EmployeeNumber: employeeNumber,
		}
		mockEmp := entityEmployee.Employee{
			EmployeeNumber: employeeNumber,
			CompanyID:      companyID,
		}
		input := &dynamodb.BatchGetItemInput{
			RequestItems: map[string]*dynamodb.KeysAndAttributes{
				tableName: {
					Keys: []map[string]*dynamodb.AttributeValue{
						mockBal.ToKey(),
						mockEmp.ToKey(),
					},
				},
			},
		}
		balAV, err := mockBal.ToItem()
		require.NoError(t, err)
		empAV, err := mockEmp.ToItem()
		require.NoError(t, err)
		output := &dynamodb.BatchGetItemOutput{
			Responses: map[string][]map[string]*dynamodb.AttributeValue{
				tableName: {
					balAV,
					empAV,
				},
			},
		}
		dClient.On("BatchGetItem", input).Return(output, nil)

		// When GetEmployeeAndBalance is called
		actEmp, actBal, actErr := repo.GetEmployeeAndBalance(employeeNumber, companyID)

		// Then employee and balance should have correct values without any error
		require.NoError(t, actErr)
		assert.NotNil(t, actEmp)
		assert.NotNil(t, actBal)
	})
}
