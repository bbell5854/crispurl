package main

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/bbell5854/crispurl/internal/repository"
)

type service struct {
	dynamoDbClient dynamodbiface.DynamoDBAPI
}

type Service interface {
	SaveShortnerRedirect(repository.ShortnerRedirect) error
}

func (svc *service) DynamoDbClient() dynamodbiface.DynamoDBAPI {
	return svc.dynamoDbClient
}

func newService(db dynamodbiface.DynamoDBAPI) *service {
	return &service{
		dynamoDbClient: db,
	}
}

func (svc *service) SaveShortnerRedirect(sr repository.ShortnerRedirect) error {
	return repository.SaveShornerRedirect(svc.dynamoDbClient, config.Env, sr)
}
