package employee

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	entityEmployee "github.com/pevin/internal-pos-service-api/domain/employee/entity"
	"github.com/rs/zerolog/log"
)

type Repo struct {
	tableName      string
	dynamodbClient dynamodbClient
	builder        employeeDynamodbBuilder
}

type RepoOpt struct {
	TableName               string
	DynamodbClient          dynamodbClient
	EmployeeDynamodbBuilder employeeDynamodbBuilder
}

type dynamodbClient interface {
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
}

type employeeDynamodbBuilder interface {
	BuildGetRFIDQueryInput(employee entityEmployee.Employee, tableName string) (*dynamodb.QueryInput, error)
}

func NewRepo(opts RepoOpt) *Repo {
	return &Repo{
		tableName:      opts.TableName,
		dynamodbClient: opts.DynamodbClient,
		builder:        opts.EmployeeDynamodbBuilder,
	}
}

func (r *Repo) GetEmployeeNumberFromRFID(rfid, companyID string) (empNumber string, err error) {
	rfidEmp := entityEmployee.Employee{
		RFID:      rfid,
		CompanyID: companyID,
	}

	queryInput, err := r.builder.BuildGetRFIDQueryInput(rfidEmp, r.tableName)
	if err != nil {
		return
	}
	resp, err := r.dynamodbClient.Query(queryInput)

	if err != nil || len(resp.Items) == 0 {
		return
	}

	emp, err := entityEmployee.FromItem(resp.Items[0])

	empNumber = emp.EmployeeNumber

	return
}

func (r *Repo) GetEmployeeAndBalance(employeeNumber, companyID string) (emp entityEmployee.Employee, bal entityEmployee.Balance, err error) {
	b := entityEmployee.Balance{
		CompanyID:      companyID,
		EmployeeNumber: employeeNumber,
	}
	e := entityEmployee.Employee{
		EmployeeNumber: employeeNumber,
		CompanyID:      companyID,
	}
	param := []map[string]*dynamodb.AttributeValue{
		b.ToKey(),
		e.ToKey(),
	}
	resp, err := r.dynamodbClient.BatchGetItem(getBatchGetItemInput(param, r.tableName))
	if err != nil {
		return
	}

	if len(resp.Responses[r.tableName]) != 2 {
		return
	}

	for _, item := range resp.Responses[r.tableName] {
		tempEmp, respErr := entityEmployee.FromItem(item)
		if respErr != nil {
			log.Error().Err(respErr).Msg("Got error unmarshalling employee from dynamodb item")
			return emp, bal, respErr
		}

		if tempEmp.IsEmployee() {
			emp = tempEmp
			continue
		}

		bal, respErr = entityEmployee.BalanceFromItem(item)
		if respErr != nil {
			log.Error().Err(respErr).Msg("Got error unmarshalling employee balance from dynamodb item")
			return emp, bal, respErr
		}
	}

	return
}

func getBatchGetItemInput(items []map[string]*dynamodb.AttributeValue, tableName string) *dynamodb.BatchGetItemInput {
	return &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			tableName: {
				Keys: items,
			},
		},
	}
}
