package main

import "github.com/aws/aws-lambda-go/lambda"

func handleLambdaEvent() (string, error) {
	return "My first assignment", nil
}

func main() {
	lambda.Start(handleLambdaEvent)
}
