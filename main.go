package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	l, _ := zap.NewProduction()
	logger = l
	defer logger.Sync() // flushes buffer, if any
}

type DefaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func myHandler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// log request info
	logger.Info("received event", zap.Any("Method", event.HTTPMethod), zap.Any("Path", event.Path), zap.Any("Body", event.Body))

	var res *events.APIGatewayProxyResponse

	if event.Path == "/hello" {

		body, _ := json.Marshal(&DefaultResponse{
			Status:  string(http.StatusOK),
			Message: "Hello World...",
		})

		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(body),
		}
	} else {
		body, _ := json.Marshal(&DefaultResponse{
			Status:  string(http.StatusOK),
			Message: "Default Path...",
		})

		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(body),
		}
	}

	return res, nil
}

func main() {
	lambda.Start(myHandler)
}
