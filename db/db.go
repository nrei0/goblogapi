package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}))
}

func CreateClient(session *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(session)
}