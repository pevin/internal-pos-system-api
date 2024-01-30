package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pevin/internal-pos-service-api/handlers/createcheckout/handler"
)

func main() {
	// init lambda handler
	lambda.StartWithOptions(handler.Handle, lambda.WithContext(context.Background()))
}
