package repository

import (
	"github.com/aws/aws-sdk-go/aws/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetConnection() *dynamodb.DynamoDB{
	sess:=session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}