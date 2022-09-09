package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AWS_SESSION *session.Session

func InitSession(region, s3Bucket, s3BucketPrefixes string) {
	AWS_SESSION = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region)},
	}))
}
