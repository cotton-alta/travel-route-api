package main

import (
	"encoding/json"
	"travel-route-api/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.PathParameters["PointId"]

	sess, err := session.NewSession()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	db := dynamodb.New(sess)

	getItem := &dynamodb.GetItemInput{
		TableName: aws.String("route-point"),
		Key: map[string]*dynamodb.AttributeValue{
			"PointId": {
				N: aws.String(path),
			},
		},
	}

	resultItem, err := db.GetItem(getItem)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 404,
		}, err
	}

	point := models.Point{}
	err = dynamodbattribute.UnmarshalMap(resultItem.Item, &point)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	responsePoint, _ := json.Marshal(point)

	return events.APIGatewayProxyResponse{
		Body:       string(responsePoint),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
