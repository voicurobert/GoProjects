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

// HandleSecondAssignmentLambda lambda funcs
func HandleSecondAssignmentLambda(ctx context.Context, data Data) (int, error) {
	return data.Income / 2, nil
}

func main() {
	lambda.Start(HandleSecondAssignmentLambda)
}
