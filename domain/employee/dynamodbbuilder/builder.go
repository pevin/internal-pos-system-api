package dynamodbbuilder

import (
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/pevin/internal-pos-service-api/domain/employee/entity"
)

type DynamodbBuilder struct{}

func NewDynamodbBuilder() *DynamodbBuilder {
	return &DynamodbBuilder{}
}

func (b *DynamodbBuilder) BuildGetRFIDQueryInput(employee entity.Employee, tableName string) (*dynamodb.QueryInput, error) {
	keyCond := expression.Key("RFIDPK").
		Equal(expression.Value(employee.GetRFIDPK())).
		And(expression.Key("RFIDSK").Equal(expression.Value(employee.GetRFIDSK())))
	builder := expression.NewBuilder().WithKeyCondition(keyCond)
	expr, err := builder.Build()
	if err != nil {
		log.Printf("Got error building Dynamodb expression: %s", err)
		return nil, err
	}

	return &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String("RFID"),
		Limit:                     aws.Int64(1),
	}, nil
}

func (b *DynamodbBuilder) BuildUpdateBalanceRequest(tableName string, bal entity.Balance, newBal float64) *dynamodb.Update {
	return &dynamodb.Update{
		TableName:        aws.String(tableName),
		Key:              bal.ToKey(),
		UpdateExpression: aws.String("SET balance = :new, updated_at = :ua"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":new": {
				N: aws.String(strconv.FormatFloat(newBal, 'f', 2, 64)),
			},
			":ua": {
				S: aws.String(bal.UpdatedAt.Format(time.RFC3339Nano)),
			},
		},
	}
}
