package apigwy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
)

type ContextHandler func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type ResponseBody struct {
	Payload interface{} `json:"payload,omitempty"`
	Message string      `json:"message,omitempty"`
}

func GetPathParameterFromRequest(req events.APIGatewayProxyRequest, paramName string) (string, error) {
	rawParamValue, ok := req.PathParameters[paramName]
	if !ok || rawParamValue == "" {
		return "", fmt.Errorf("parameter not found: %s", paramName)
	}

	paramValue, err := url.QueryUnescape(rawParamValue)
	if err != nil {
		return "", err
	}

	return paramValue, nil
}

func CreateApiGatewayProxyResponse(statusCode int, body ResponseBody) events.APIGatewayProxyResponse {
	headers := map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"}
	bodyString, _ := json.Marshal(body)

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       string(bodyString),
	}
}

func CreateApiGatewayProxyRedirectResponse(redirectUrl string) events.APIGatewayProxyResponse {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Location":     redirectUrl,
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMovedPermanently,
		Headers:    headers,
	}
}
