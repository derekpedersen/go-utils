package aws

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var DYNAMO_DB *dynamodb.DynamoDB

func InitDynamoDb() {
	DYNAMO_DB = dynamodb.New(AWS_SESSION)
	logrus.WithField("DYNAMO_DB", *DYNAMO_DB).Info("Configured Dynamo DB Connection")
}

func DynamoDbBatchWrite(batch *dynamodb.BatchWriteItemInput, logid uuid.UUID) (*dynamodb.BatchWriteItemOutput, error) {
	if DYNAMO_DB == nil {
		return nil, os.ErrInvalid
	}
	return DynamoDbBatchWriteWithClient(DYNAMO_DB, batch, logid)
}

func DynamoDbBatchWriteWithClient(client dynamodbiface.DynamoDBAPI, batch *dynamodb.BatchWriteItemInput, logid uuid.UUID) (*dynamodb.BatchWriteItemOutput, error) {
	if client == nil {
		return nil, os.ErrInvalid
	}
	output, err := client.BatchWriteItem(batch)
	if err != nil {
		text, _ := json.Marshal(map[string]interface{}{
			"Batch":  batch,
			"Output": output,
		})
		if mkdirErr := os.MkdirAll(".logs/debug", 0755); mkdirErr != nil {
			logrus.WithError(mkdirErr).Warn("failed to create DynamoDB debug directory")
		} else if writeErr := os.WriteFile(filepath.Join(".logs/debug", fmt.Sprint(time.Now().Unix())+"_failure_batchwrite_"+logid.String()+".json"), text, 0644); writeErr != nil {
			logrus.WithError(writeErr).Warn("failed to write DynamoDB debug file")
		}
	}
	return output, err
}
