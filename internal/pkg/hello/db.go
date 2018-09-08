package hello

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-2"))

func getItem(isbn string)(*book, error) {
	input := dynamodb.GetItemInput {
		TableName: aws.String("Books"),
		Key: dynamodb.AttributeValue {
			"ISBN" : {S: aws.String(isbn),}
		},
	}
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	bk := new(book)
	err = dynamodbattribute.UnmarshalMap(result.Item, bk)
	if (err != nil) {
		return nil, err
	}

	return bk, nil
}

