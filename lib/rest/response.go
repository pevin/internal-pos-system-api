package rest

import (
	"encoding/json"
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

type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
}

func DefaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func BadRequestResponse(msg string) Response {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: false,
	}
	return Response{
		StatusCode: 400,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func NotFoundResponse(msg string) Response {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: false,
	}
	return Response{
		StatusCode: 404,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func UnauthorizedResponse(msg string) Response {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: false,
	}
	return Response{
		StatusCode: 401,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func OkResponse(data interface{}, msg string) Response {
	rb := ResponseBody{
		Data:    data,
		Message: msg,
		Success: true,
	}
	return Response{
		StatusCode: 200,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}

func EmptyOkResponse(msg string) Response {
	rb := ResponseBody{
		Data:    struct{}{},
		Message: msg,
		Success: true,
	}
	return Response{
		StatusCode: 200,
		Headers:    DefaultHeaders(),
		Body:       rb.String(),
	}
}
