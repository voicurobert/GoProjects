package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

//Data struct
type Data struct {
	Age    int `json:"age"`
	Height int `json:"height"`
	Income int `json:"income"`
}

// HandleDeleteDataLambda lambda funcs
func HandleDeleteDataLambda(ctx context.Context, data Data) (string, error) {
	return "Deleted", nil
}

func main() {
	lambda.Start(HandleDeleteDataLambda)
}
