package aws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var DYNAMO_DB *dynamodb.DynamoDB

func InitDynamoDb() {
	DYNAMO_DB = dynamodb.New(AWS_SESSION)
	logrus.WithField("DYNAMO_DB", *DYNAMO_DB).Info("Configured Dynamo DB Connection")
}

func DynamoDbBatchWrite(batch *dynamodb.BatchWriteItemInput, logid uuid.UUID) (*dynamodb.BatchWriteItemOutput, error) {
	output, err := DYNAMO_DB.BatchWriteItem(batch)
	if err != nil {
		text, _ := json.Marshal(map[string]interface{}{
			"Batch":  batch,
			"Output": output,
		})
		err = ioutil.WriteFile(".logs/debug/"+fmt.Sprint(time.Now().Unix())+"_failure_batchwrite"+"_"+strconv.Itoa(rand.Int())+"_"+".json", text, 0644)
		if err != nil {
			logrus.Error(err)
		}
	}
	return output, nil
}
