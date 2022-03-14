package repository

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var shortnerRedirectsTableName = "shortner-redirects"
var shortnerRedirectsPartitionKey = "shortnerID"

var ErrShortnerRedirectNotFound = errors.New("shortner-redirect not found")

type ShortnerRedirect struct {
	ShortnerID  string `dynamodbav:"shortnerID"`
	RedirectURL string `dynamodbav:"redirectURL"`
}

func GetShortnerRedirect(db dynamodbiface.DynamoDBAPI, env, shortnerId string) (*ShortnerRedirect, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(buildTableName(shortnerRedirectsTableName, env)),
		Key: map[string]*dynamodb.AttributeValue{
			shortnerRedirectsPartitionKey: {
				S: aws.String(shortnerId),
			},
		},
	}

	res, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}

	// add custom error message for record not found
	if res.Item == nil {
		return nil, ErrShortnerRedirectNotFound
	}

	sr := &ShortnerRedirect{}
	err = dynamodbattribute.UnmarshalMap(res.Item, &sr)
	if err != nil {
		return nil, err
	}

	return sr, nil
}

func SaveShornerRedirect(db dynamodbiface.DynamoDBAPI, env string, sr ShortnerRedirect) error {
	av, err := dynamodbattribute.MarshalMap(sr)
	if err != nil {
		return errors.New("unable to marshal shortnerRedirect")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(buildTableName(shortnerRedirectsTableName, env)),
	}

	if _, err = db.PutItem(input); err != nil {
		if reqerr, ok := err.(awserr.RequestFailure); ok {
			return fmt.Errorf("request failed %s %s %s", reqerr.Code(), reqerr.Message(), reqerr.RequestID())
		}

		return err
	}

	return nil
}
