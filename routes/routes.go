package routes

import (
	"context"
	"fmt"
	"go-sqlc/controllers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func SetupLambda() {
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse

	switch request.HTTPMethod {
	case "GET":
		switch request.RequestContext.Path {
		case "/get-customer-orders/{id}":
			fmt.Println("Got Here Get Customer Orders")
			return controllers.GetCustomerOrders(ctx, request)
		case "/hello-world":
			return controllers.HelloWorld()
		}
	case "POST":
		switch request.RequestContext.Path {
		case "/create-customer":
			return controllers.CreateCustomer(ctx, request)
		}
	case "PUT":
		switch request.RequestContext.Path {
		case "/update-customer":
			return controllers.UpdateCustomer(ctx, request)
		}
	case "DELETE":
		switch request.RequestContext.Path {
		case "/hard-delete-customer/{id}":
			return controllers.HardDeleteCustomer(ctx, request)
		case "/soft-delete-customer/{id}":
			return controllers.SoftDeleteCustomer(ctx, request)
		}
	default:
		response = events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found Path.",
		}
	}

	return &response, nil
}
