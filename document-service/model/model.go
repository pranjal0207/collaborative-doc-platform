package model

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DocumentModel struct {
	DynamoDB  *dynamodb.Client
	TableName string
}

