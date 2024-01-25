package entity

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	userEntity "github.com/pevin/internal-pos-service-api/domain/user/entity"
)

type Employee struct {
	PK             string                 `dynamodbav:"PK" json:"-"`
	SK             string                 `dynamodbav:"SK" json:"-"`
	GSI1PK         string                 `dynamodbav:"GSI1PK" json:"-"`
	GSI1SK         string                 `dynamodbav:"GSI1SK" json:"-"`
	RFID_PK        string                 `dynamodbav:"RFIDPK,omitempty" json:"-"`
	RFID_SK        string                 `dynamodbav:"RFIDSK,omitempty" json:"-"`
	EmployeeNumber string                 `dynamodbav:"employee_number" json:"employee_number"`
	CompanyID      string                 `dynamodbav:"company_id" json:"company_id"`
	UserID         string                 `dynamodbav:"user_id,omitempty" json:"user_id,omitempty"`
	FirstName      string                 `dynamodbav:"first_name" json:"first_name"`
	LastName       string                 `dynamodbav:"last_name" json:"last_name"`
	ImageUrl       string                 `dynamodbav:"image_url,omitempty" json:"image_url"`
	BirthDate      string                 `dynamodbav:"birth_date" json:"birth_date"`
	HiredDate      string                 `dynamodbav:"hired_date" json:"hired_date"`
	RFID           string                 `dynamodbav:"rfid" json:"rfid"`
	CreditLimit    float64                `dynamodbav:"credit_limit,omitempty" json:"credit_limit"`
	Discount       float64                `dynamodbav:"discount,omitempty" json:"discount"`
	Status         string                 `dynamodbav:"status" json:"status"`
	CreatedAt      time.Time              `dynamodbav:"created_at,omitempty" json:"created_at,omitempty"`
	CreatedBy      userEntity.PartialUser `dynamodbav:"created_by" json:"created_by,omitempty"`
	UpdatedAt      time.Time              `dynamodbav:"updated_at,omitempty" json:"updated_at"`
	UpdatedBy      userEntity.PartialUser `dynamodbav:"updated_by" json:"updated_by"`
}

// ToItem returns the dynamodb item for the employee
func (e *Employee) ToItem() (map[string]*dynamodb.AttributeValue, error) {
	e.genTableKey()
	e.genGSI1Key()
	e.genRFIDIndexKey()
	av, err := dynamodbattribute.MarshalMap(e)

	return av, err
}

// FromItem unmarshal the dynamodb item to the employee struct
func FromItem(item map[string]*dynamodb.AttributeValue) (Employee, error) {
	e := Employee{}
	err := dynamodbattribute.UnmarshalMap(item, &e)
	if err != nil {
		return e, err
	}
	return e, nil
}

// ToKey returns the key for the employee
func (e *Employee) ToKey() map[string]*dynamodb.AttributeValue {
	e.genTableKey()
	key := map[string]*dynamodb.AttributeValue{
		"PK": {
			S: aws.String(e.PK),
		},
		"SK": {
			S: aws.String(e.SK),
		},
	}
	return key
}

func (e *Employee) GetRFIDPK() string {
	e.genRFIDPK()
	return e.RFID_PK
}

func (e *Employee) GetRFIDSK() string {
	e.genRFIDSK()
	return e.RFID_SK
}

func (e *Employee) IsEmployee() bool {
	return e.SK[:strings.IndexByte(e.SK, '#')] == "EMPLOYEE"
}

func (e *Employee) GetEmployeeName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) genTableKey() {
	e.PK = "EMPLOYEE#" + e.EmployeeNumber
	e.SK = "EMPLOYEE#" + e.EmployeeNumber
}

func (e *Employee) genGSI1Key() {
	e.GSI1PK = "COMPANY#" + e.CompanyID
	e.GSI1SK = "EMPLOYEE#" + e.CreatedAt.Format(time.RFC3339) + "#" + e.EmployeeNumber // Added created_at in the SK to sort by latest
}

func (e *Employee) genRFIDIndexKey() {
	e.genRFIDPK()
	e.genRFIDSK()
}

func (e *Employee) genRFIDPK() {
	if e.RFID_PK == "" {
		e.RFID_PK = "COMPANY#" + e.CompanyID
	}
}

func (e *Employee) genRFIDSK() {
	if e.RFID_SK == "" {
		e.RFID_SK = "RFID#" + e.RFID
	}
}
