package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Request struct
type Request struct {
}

// Response struct
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// PersonData struct
type PersonData struct {
	Data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"personData"`
}

// Data struct
type Data struct {
	Age    int `json:"age"`
	Height int `json:"height"`
	Income int `json:"income"`
}

// HandleLambdaEvent func
func HandleLambdaEvent(ctx context.Context, data Data) (Response, error) {
	/*
		lc, _ := lambdacontext.FromContext(ctx)
		fmt.Println(lc)

		resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
		resp.StatusCode = 200
		resp.Headers["Control-Access-Allow-Origin"] = "*"

		resp.Body = string("My Lambda")
	*/
	return Response{Message: fmt.Sprintf("Age %d", data.Age*2)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
