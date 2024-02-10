package main

import (
	"context"
	c "scheduler-api/config"
	db "scheduler-api/db"
	r "scheduler-api/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.LambdaFunctionURLRequest) (Response, error) {
	return Response{Body: "It works!", StatusCode: 200}, nil
}

func main() {
	c.EnvSetup()
	db.Connect()

	e := echo.New()

	r.InitRoutes(e)
	e.Logger.Fatal(e.Start(":3500"))
	lambda.Start(Handler)
}
