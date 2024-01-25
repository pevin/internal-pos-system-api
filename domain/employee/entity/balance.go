package entity

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Balance struct {
	PK             string    `dynamodbav:"PK" json:"-"`
	SK             string    `dynamodbav:"SK" json:"-"`
	GSIPK1         string    `dynamodbav:"GSI1PK" json:"-"`
	GSISK1         string    `dynamodbav:"GSI1SK" json:"-"`
	EmployeeNumber string    `dynamodbav:"employee_number" json:"employee_number"`
	CompanyID      string    `dynamodbav:"company_id" json:"company_id"`
	Balance        float64   `dynamodbav:"balance" json:"balance"`
	CreditLimit    float64   `dynamodbav:"credit_limit" json:"credit_limit"`
	CreatedAt      time.Time `dynamodbav:"created_at" json:"created_at"`
	UpdatedAt      time.Time `dynamodbav:"updated_at" json:"updated_at"`
	ReplenishedAt  string    `dynamodbav:"replenished_at" json:"replenished_at"`

	// Add discount here?
}

func BalanceFromItem(item map[string]*dynamodb.AttributeValue) (Balance, error) {
	b := Balance{}
	err := dynamodbattribute.UnmarshalMap(item, &b)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (b *Balance) ToItem() (map[string]*dynamodb.AttributeValue, error) {
	b.genTableKey()
	b.genGSI1Key()
	item, err := dynamodbattribute.MarshalMap(b)

	return item, err
}

func (b *Balance) ToKey() map[string]*dynamodb.AttributeValue {
	b.genTableKey()
	key := map[string]*dynamodb.AttributeValue{
		"PK": {
			S: aws.String(b.PK),
		},
		"SK": {
			S: aws.String(b.SK),
		},
	}
	return key
}

func (b *Balance) genTableKey() {
	b.PK = "EMPLOYEE#" + b.EmployeeNumber
	b.SK = "BALANCE#" + b.EmployeeNumber
}

func (b *Balance) genGSI1Key() {
	b.GSIPK1 = "COMPANY#" + b.CompanyID
	b.GSISK1 = "BALANCE#" + b.EmployeeNumber
}
