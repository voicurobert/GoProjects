package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//Data struct
type Data struct {
	Type string `json:"myType"`
}

// Item struct
type Item struct {
	UserID string
	Age    int
	Height int
	Income int
}

//HandleGetDataLambda func
func HandleGetDataLambda(context context.Context, data Data) ([]Item, error) {
	t := data.Type
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	if t == "all" {
		return getAllItems(svc, "compare-yourself")
	} else if t == "single" {
		return getSingleItem(svc, "compare-yourself")
	} else {
		return nil, fmt.Errorf("Wrong parameter %s", t)
	}
}

func getAllItems(dynamo *dynamodb.DynamoDB, tableName string) ([]Item, error) {
	params := &dynamodb.ScanInput{TableName: aws.String(tableName)}
	result, err := dynamo.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	items := make([]Item, int(*result.Count), int(*result.Count))
	for idx, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err == nil {
			items[idx] = item
		}
	}
	return items, nil
}

func getSingleItem(dynamo *dynamodb.DynamoDB, tableName string) ([]Item, error) {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String("asdahdsajk34343"),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Error retrieving item %s", err)
	}
	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling item %s", err)
	}
	items := make([]Item, 1, 1)
	items[0] = item
	return items, nil
}

func main() {
	lambda.Start(HandleGetDataLambda)
}
