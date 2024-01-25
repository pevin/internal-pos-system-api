package entity

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	userEntity "github.com/pevin/internal-pos-service-api/domain/user/entity"
)

type Checkout struct {
	PK               string                 `dynamodbav:"PK" json:"-"` // PK: EMPLOYEE#<ID>
	SK               string                 `dynamodbav:"SK" json:"-"` // SK: CHECKOUT#YYYYMMDDHHMMSSMS
	GSI1PK           string                 `dynamodbav:"GSI1PK" json:"-"`
	GSI1SK           string                 `dynamodbav:"GSI1SK" json:"-"`
	GSI2PK           string                 `dynamodbav:"GSI2PK" json:"-"`
	GSI2SK           string                 `dynamodbav:"GSI2SK" json:"-"`
	ID               string                 `dynamodbav:"checkout_id" json:"checkout_id"`
	CompanyID        string                 `dynamodbav:"com_id" json:"company_id"`
	ConcessionaireID string                 `dynamodbav:"con_id" json:"concessionaire_id"`
	EmployeeNumber   string                 `dynamodbav:"emp_num" json:"employee_number"`
	DiscountAmount   float64                `dynamodbav:"disc_amt" json:"discount_amount"`
	TotalGrossAmount float64                `dynamodbav:"total_gamt" json:"total_gross_amount"`
	TotalNetAmount   float64                `dynamodbav:"total_namt" json:"total_net_amount"`
	TotalCalories    float64                `dynamodbav:"total_cal" json:"total_calories"`
	Station          string                 `dynamodbav:"station" json:"station"`
	CreatedAt        time.Time              `dynamodbav:"created_at" json:"created_at"`
	CreatedBy        userEntity.PartialUser `dynamodbav:"created_by" json:"created_by"`
	IsVoid           bool                   `dynamodbav:"voided" json:"is_void"`
	VoidedAt         time.Time              `dynamodbav:"voided_at" json:"voided_at"`
	VoidedBy         userEntity.PartialUser `dynamodbav:"voided_by" json:"voided_by"`
	CheckoutItems    []CheckoutItem         `dynamodbav:"items" json:"items"`
}

type CheckoutItem struct {
	MealCode         string  `dynamodbav:"code" json:"meal_code"`
	MealName         string  `dynamodbav:"name" json:"meal_name"`
	Category         string  `dynamodbav:"cat" json:"category"`
	Price            float64 `dynamodbav:"price" json:"price"`
	Calories         float64 `dynamodbav:"cal" json:"calories"`
	Quantity         int     `dynamodbav:"qty" json:"qty"`
	SubTotalPrice    float64 `dynamodbav:"total_p" json:"sub_total_price"`
	SubTotalCalories float64 `dynamodbav:"total_c" json:"sub_total_calories"`
}

func (co *Checkout) ToItem() (map[string]*dynamodb.AttributeValue, error) {
	co.genTableKey()
	co.genGSI1Key()
	co.genGSI2Key()
	av, err := dynamodbattribute.MarshalMap(co)

	return av, err
}

func (co *Checkout) ToKey() map[string]*dynamodb.AttributeValue {
	co.genTableKey()
	key := map[string]*dynamodb.AttributeValue{
		"PK": {
			S: aws.String(co.PK),
		},
		"SK": {
			S: aws.String(co.SK),
		},
	}
	return key
}

func (co *Checkout) genTableKey() {
	if co.PK == "" {
		co.PK = "EMPLOYEE#" + co.EmployeeNumber
	}
	if co.SK == "" {
		co.SK = "CHECKOUT#" + co.ID
	}
}

func (co *Checkout) genGSI1Key() {
	if co.GSI1PK == "" {
		co.GSI1PK = "COMPANY#" + co.CompanyID
	}
	if co.GSI1SK == "" {
		co.GSI1SK = "CHECKOUT#" + co.ID + "#" + co.EmployeeNumber
	}
}

func (co *Checkout) genGSI2Key() {
	if co.GSI2PK == "" {
		co.GSI2PK = "CONCESSIONAIRE#" + co.ConcessionaireID
	}
	if co.GSI2SK == "" {
		co.GSI2SK = "CHECKOUT#" + co.ID + "#" + co.EmployeeNumber
	}
}
