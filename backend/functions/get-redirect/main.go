package main

import (
	log "github.com/sirupsen/logrus"

	env "github.com/Netflix/go-env"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/bbell5854/crispurl/internal/utils/logger"
)

var config struct {
	Env       string `env:"ENV"`
	AwsRegion string `env:"REGION"`
	LogLevel  string `env:"LOG_LEVEL"`
}

func main() {
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		log.Fatal("Unable to unmarshal environment variables", err)
	}

	logger.Setup(config.LogLevel)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamodb.New(sess, &aws.Config{
		Region: aws.String(config.AwsRegion),
	})

	svc := newService(db)

	lambda.Start(handlerRouter(svc))
}
