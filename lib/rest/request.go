package rest

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	entityUser "github.com/pevin/internal-pos-service-api/domain/user/entity"
)

type Request struct {
	User *entityUser.User
	Body string
}

func FromAPIGatewayProxy(req events.APIGatewayProxyRequest) (Request, error) {
	user, err := getUserFromAPIGatewayProxy(req.RequestContext)
	if err != nil {
		return Request{}, err
	}
	return Request{
		Body: req.Body,
		User: user,
	}, nil
}

func getUserFromAPIGatewayProxy(reqContext events.APIGatewayProxyRequestContext) (u *entityUser.User, err error) {
	authClaim, err := json.Marshal(reqContext.Authorizer["claims"])
	if err != nil {
		log.Printf("Got error marshaling auth claims: %s", err)
		return
	}

	err = json.Unmarshal(authClaim, &u)
	if err != nil {
		log.Printf("Got error unmarshaling auth claim: %s", err)
		return
	}

	return u, nil
}
