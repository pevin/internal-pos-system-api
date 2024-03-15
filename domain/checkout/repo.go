package checkout

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pevin/internal-pos-service-api/domain/checkout/entity"
	entityEmployee "github.com/pevin/internal-pos-service-api/domain/employee/entity"
)

type Repo struct {
	tableName               string
	dynamodbClient          CheckoutDynamodbClient
	employeeDynamodbBuilder CheckoutEmployeeDynamodbBuilder
}

type RepoOpt struct {
	TableName               string
	DynamodbClient          CheckoutDynamodbClient
	EmployeeDynamodbBuilder CheckoutEmployeeDynamodbBuilder
}

//go:generate mockery --name CheckoutDynamodbClient
type CheckoutDynamodbClient interface {
	TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error)
}

//go:generate mockery --name CheckoutEmployeeDynamodbBuilder
type CheckoutEmployeeDynamodbBuilder interface {
	BuildUpdateBalanceRequest(tableName string, bal entityEmployee.Balance, newBal float64) *dynamodb.Update
}

func NewRepo(opts RepoOpt) *Repo {
	return &Repo{
		dynamodbClient:          opts.DynamodbClient,
		employeeDynamodbBuilder: opts.EmployeeDynamodbBuilder,
		tableName:               opts.TableName,
	}
}

// Transact will insert the Checkout entity and update the Balance entity in dynamodb
func (r *Repo) Transact(co entity.Checkout, bal entityEmployee.Balance, newBal float64) (err error) {
	av, err := co.ToItem()
	if err != nil {
		return
	}
	_, err = r.dynamodbClient.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					Item:      av,
					TableName: aws.String(r.tableName),
				},
			},
			{
				Update: r.employeeDynamodbBuilder.BuildUpdateBalanceRequest(r.tableName, bal, newBal),
			},
		},
	})
	return
}
