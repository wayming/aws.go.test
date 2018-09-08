package awsctl

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-2"))

func randString(len int) string {

	bytes := make([]byte, len)

	for i := range bytes {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}

// Load random data into dynamondb
func Load(numberOfRows int) {

	for i := 1; i <= numberOfRows; i++ {

		input := &dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"ISBN": {
					S: aws.String(strconv.Itoa(100+rand.Intn(899)) + "-" +
						strconv.Itoa(100000000+rand.Intn(899999999))),
				},
				"Title": {
					S: aws.String(randString(4) + " " + randString(3) + " " + randString(2)),
				},
				"Author": {
					S: aws.String(randString(8)),
				},
			},
			TableName: aws.String("Books"),
		}
		fmt.Println("Inserting " + strconv.Itoa(i) + "th record")

		result, err := db.PutItem(input)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				fmt.Println("Error:", awsErr.Code(), awsErr.Message())
			}
		} else {
			fmt.Println(result)
		}
	}
	return

}
