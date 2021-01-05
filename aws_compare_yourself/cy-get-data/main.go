package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

//Data struct
type Data struct {
	Type string `json:"myType"`
}

//HandleGetDataLambda func
func HandleGetDataLambda(context context.Context, data Data) (string, error) {

	t := data.Type

	if t == "all" {
		return "All the data", nil
	} else if t == "single" {
		return "Single data", nil
	} else {
		return "", fmt.Errorf("Wrong parameter %s", t)
	}
}

func main() {
	lambda.Start(HandleGetDataLambda)
}
