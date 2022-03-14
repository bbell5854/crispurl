package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bbell5854/crispurl/internal/platform/apigwy"
	"github.com/bbell5854/crispurl/internal/repository"
	log "github.com/sirupsen/logrus"
)

func handlerRouter(svc Service) apigwy.ContextHandler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (e events.APIGatewayProxyResponse, err error) {
		var statusCode int

		shortnerID, found := req.PathParameters["shortnerID"]
		if !found {
			log.Debug("shortnerID not found in path parameters")
			statusCode = http.StatusNotFound
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode)}), nil
		}

		parsedShortnerID, err := url.QueryUnescape(shortnerID)
		if err != nil {
			log.Debugf("Unable to escape shortnerID: %s", err.Error())
			statusCode = http.StatusBadRequest
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode)}), nil
		}

		redirectUrl, err := svc.GetRedirectUrl(parsedShortnerID)
		if err == repository.ErrShortnerRedirectNotFound {
			log.Debugf("redirect not found for shortnerID: %s", parsedShortnerID)
			statusCode = http.StatusNotFound
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode)}), nil
		}

		if err != nil {
			log.Errorf("unable to fetch redirect - %s : %s", parsedShortnerID, err.Error())
			statusCode = http.StatusInternalServerError
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode)}), nil
		}

		log.Infof("redirect found - %s : %s", parsedShortnerID, redirectUrl)
		return apigwy.CreateApiGatewayProxyRedirectResponse(redirectUrl), nil
	}
}
