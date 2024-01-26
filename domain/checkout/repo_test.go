package checkout_test

import (
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pevin/internal-pos-service-api/domain/checkout"
	entityCheckout "github.com/pevin/internal-pos-service-api/domain/checkout/entity"
	"github.com/pevin/internal-pos-service-api/domain/employee/entity"
	empEntity "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	entityEmployee "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockedDynamodbClient struct {
	mock.Mock
}

func (mr *mockedDynamodbClient) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	args := mr.Called(input)
	return args.Get(0).(*dynamodb.TransactWriteItemsOutput), args.Error(1)
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

func (medb *mockEmployeeDynamodbBuilder) BuildGetRFIDQueryInput(employee empEntity.Employee, tableName string) (*dynamodb.QueryInput, error) {
	args := medb.Called(employee, tableName)
	return args.Get(0).(*dynamodb.QueryInput), args.Error(1)
}
func (medb *mockEmployeeDynamodbBuilder) BuildUpdateBalanceRequest(tableName string, bal entity.Balance, newBal float64) *dynamodb.Update {
	args := medb.Called(tableName, bal, newBal)
	return args.Get(0).(*dynamodb.Update)
}

func TestTransactSuccess(t *testing.T) {
	// Given a valid dynamodb command
	dynamodbClient := new(mockedDynamodbClient)
	builder := new(mockEmployeeDynamodbBuilder)
	tableName := "test-table-name"
	opts := checkout.RepoOpt{
		TableName:               tableName,
		DynamodbClient:          dynamodbClient,
		EmployeeDynamodbBuilder: builder,
	}

	co := entityCheckout.Checkout{}
	av, err := co.ToItem()
	require.NoError(t, err)
	bal := entityEmployee.Balance{}
	balAfter := 2500.00

	dynamodbUpdateRequest := &dynamodb.Update{ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
		":new": {
			N: aws.String(strconv.FormatFloat(balAfter, 'f', 2, 64)),
		},
	}}
	builder.On("BuildUpdateBalanceRequest", tableName, bal, balAfter).Return(dynamodbUpdateRequest)

	trxInput := &dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					Item:      av,
					TableName: aws.String(tableName),
				},
			},
			{
				Update: dynamodbUpdateRequest,
			},
		},
	}
	dynamodbClient.On("TransactWriteItems", trxInput).Return(&dynamodb.TransactWriteItemsOutput{}, nil)
	repo := checkout.NewRepo(opts)

	// When transact is initiated
	transactErr := repo.Transact(co, bal, balAfter)

	// Then there should be no error
	assert.NoError(t, transactErr)
}
