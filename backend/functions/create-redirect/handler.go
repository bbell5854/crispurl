package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/bbell5854/crispurl/internal/platform/apigwy"
	"github.com/bbell5854/crispurl/internal/repository"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type requestBody struct {
	RedirectUrl string `json:"redirectUrl"`
}

type responseBody struct {
	ShortUrl    string `json:"shortUrl"`
	RedirectURL string `json:"redirectUrl"`
}

func handlerRouter(svc Service) apigwy.ContextHandler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (e events.APIGatewayProxyResponse, err error) {
		var statusCode int

		var b requestBody
		if err := json.Unmarshal([]byte(req.Body), &b); err != nil {
			log.Debugf("unable to unmarshal request body: %s", err.Error())
			statusCode = http.StatusBadRequest
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode)}), nil
		}

		if b.RedirectUrl == "" {
			log.Debug("redirectUrl not found in request")
			statusCode = http.StatusUnprocessableEntity
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode)}), nil
		}

		if !govalidator.IsURL(b.RedirectUrl) {
			log.Debugf("redirectUrl not a valid url: %s", b.RedirectUrl)
			statusCode = http.StatusUnprocessableEntity
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: "redirectURL provided is not a valid url"}), nil
		}

		// Do not allow recursive redirects
		if strings.Contains(b.RedirectUrl, config.ApiBaseUrl) {
			log.Errorf("potential recursive redirect attempted: %s", b.RedirectUrl)
			statusCode = http.StatusUnprocessableEntity
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: "redirectURL provided is not a valid url"}), nil
		}

		sr := repository.ShortnerRedirect{
			ShortnerID:  uuid.NewString()[0:8],
			RedirectURL: addUrlScheme(b.RedirectUrl),
		}

		if err := svc.SaveShortnerRedirect(sr); err != nil {
			log.Errorf("error saving shortnerredirect to dynamodb: %s", err.Error())
			statusCode = http.StatusInternalServerError
			return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode)}), nil
		}

		resp := responseBody{
			ShortUrl:    fmt.Sprintf("%s/%s", config.ApiBaseUrl, sr.ShortnerID),
			RedirectURL: sr.RedirectURL,
		}

		log.Infof("shortner redirect created: %s - %s", resp.ShortUrl, resp.RedirectURL)
		statusCode = http.StatusOK
		return apigwy.CreateApiGatewayProxyResponse(statusCode, apigwy.ResponseBody{Message: http.StatusText(statusCode), Payload: resp}), nil
	}
}

func addUrlScheme(redirectUrl string) string {
	parsedUrl, _ := url.Parse(redirectUrl)
	if parsedUrl.Scheme == "" {
		return fmt.Sprintf("http://%s", redirectUrl)
	}

	return redirectUrl
}
