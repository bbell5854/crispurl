package main

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/bbell5854/crispurl/internal/repository"
)

type service struct {
	dynamoDbClient dynamodbiface.DynamoDBAPI
}

type Service interface {
	GetRedirectUrl(string) (string, error)
}

func (svc *service) DynamoDbClient() dynamodbiface.DynamoDBAPI {
	return svc.dynamoDbClient
}

func newService(db dynamodbiface.DynamoDBAPI) *service {
	return &service{
		dynamoDbClient: db,
	}
}

func (svc *service) GetRedirectUrl(urlID string) (string, error) {
	sr, err := repository.GetShortnerRedirect(svc.dynamoDbClient, config.Env, urlID)
	if err != nil {
		return "", err
	}

	return sr.RedirectURL, nil
}
