package main

import (
	"flag"

	"aws.go.test/internal/pkg/awsctl"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-2"))

func main() {
	loadOption := flag.Int("load", 0, "Load # of rows with random data into dynamodb")

	flag.Parse()

	if *loadOption > 0 {
		awsctl.Load(*loadOption)
	}

}
