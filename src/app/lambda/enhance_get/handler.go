package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/conamu/go-cdk-fullstack-web-app-template/src/pkg/constants"
	"github.com/conamu/go-cdk-fullstack-web-app-template/src/pkg/domain/enhance"
	"github.com/conamu/go-cdk-fullstack-web-app-template/src/pkg/logger"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	log := logger.Create()
	ctx = context.WithValue(ctx, constants.CTX_LOGGER, log)

	log.Info("Enhance endpoint hit!")

	name := req.QueryStringParameters["name"]

	log.Info("Name hit with " + name)

	body := enhance.Enhance(name)

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              "<p>" + body.Name + "</p>",
		IsBase64Encoded:   false,
	}, nil
}
