package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Response struct
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Data struct
type Data struct {
	Age    int `json:"age"`
	Height int `json:"height"`
	Income int `json:"income"`
}

// Item struct to hold dynamo db items
type Item struct {
	UserID string
	Age    int
	Height int
	Income int
}

// StoreDataLambda func
func StoreDataLambda(ctx context.Context, data Data) (Response, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	userID := "user_" + strconv.Itoa(rand.Int())

	item := Item{UserID: userID, Age: data.Age, Height: data.Height, Income: data.Income}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tableName := "compare-yourself"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return Response{Message: "Added item to dynamo db " + userID}, nil
}

func main() {
	lambda.Start(StoreDataLambda)
}
