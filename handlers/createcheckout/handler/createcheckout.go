package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pevin/internal-pos-service-api/domain/checkout"
	"github.com/pevin/internal-pos-service-api/domain/employee"
	"github.com/pevin/internal-pos-service-api/domain/employee/dynamodbbuilder"
)

func Handle(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	// init checkout service
	tableName := os.Getenv("TABLE_NAME")
	awsRegion := os.Getenv("APP_AWS_REGION")
	conf := &aws.Config{Region: aws.String(awsRegion)}
	sess, err := session.NewSession(conf)
	if err != nil {
		fmt.Println("Error creating aws session: ", err)
		panic(err)
	}
	dynamodbClient := dynamodb.New(sess)

	empBuilder := dynamodbbuilder.NewDynamodbBuilder()

	empRepo := employee.NewRepo(employee.RepoOpt{
		TableName:               tableName,
		EmployeeDynamodbBuilder: empBuilder,
		DynamodbClient:          dynamodbClient,
	})

	checkoutRepo := checkout.NewRepo(checkout.RepoOpt{
		TableName:               tableName,
		DynamodbClient:          dynamodbClient,
		EmployeeDynamodbBuilder: empBuilder,
	})

	checkoutService := checkout.NewService(checkout.ServiceOpt{
		EmployeeRepo: empRepo,
		CheckoutRepo: checkoutRepo,
	})

	res, err = checkoutService.Create(ctx, req)
	return
}
