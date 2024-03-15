package rest

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type ResponseBody struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func (rb *ResponseBody) String() string {
	res, err := json.Marshal(&rb)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func DefaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func BadRequestResponse(msg string) events.APIGatewayProxyResponse {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: false,
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func NotFoundResponse(msg string) events.APIGatewayProxyResponse {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: false,
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func UnauthorizedResponse(msg string) events.APIGatewayProxyResponse {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: false,
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 401,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func OkResponse(data interface{}, msg string) events.APIGatewayProxyResponse {
	rb := ResponseBody{
		Data:    data,
		Message: msg,
		Success: true,
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func EmptyOkResponse(msg string) events.APIGatewayProxyResponse {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: true,
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}
