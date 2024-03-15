package checkout_test

import (
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pevin/internal-pos-service-api/domain/checkout"
	entityCheckout "github.com/pevin/internal-pos-service-api/domain/checkout/entity"
	"github.com/pevin/internal-pos-service-api/domain/checkout/mocks"
	entityEmployee "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactSuccess(t *testing.T) {
	// Given a valid dynamodb command
	dynamodbClient := new(mocks.CheckoutDynamodbClient)
	builder := new(mocks.CheckoutEmployeeDynamodbBuilder)
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
