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

// StoreDataLambda func
func StoreDataLambda(ctx context.Context, data Data) (Response, error) {
	return Response{Message: fmt.Sprintf("Age %d", data.Age*2)}, nil
}

func main() {
	lambda.Start(StoreDataLambda)
}
