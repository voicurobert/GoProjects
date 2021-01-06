package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func customAuthLambda(ctx context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := request.AuthorizationToken
	if token == "allow" {
		return genPolicy("adasdsadasdadas3", "allow", request.MethodArn), nil
	} else if token == "deny" {
		return genPolicy("adasdsadasdadas3", "deny", request.MethodArn), nil
	}
	return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Error")
}

func genPolicy(principalID, effect string, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}
	policy := new(events.APIGatewayCustomAuthorizerPolicy)
	policy.Version = "2012-10-17"
	policy.Statement = []events.IAMPolicyStatement{
		{
			Action:   []string{"execute-api:Invoke"},
			Effect:   effect,
			Resource: []string{resource},
		},
	}
	authResponse.PolicyDocument = *policy

	return authResponse
}

func main() {
	lambda.Start(customAuthLambda)
}
