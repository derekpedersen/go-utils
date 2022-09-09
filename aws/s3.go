package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
)

func Download(filename, s3bucket string) error {
	file, err := os.Create(filename)
	if err != nil {
		logrus.Error(err)
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(AWS_SESSION)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s3bucket),
			Key:    aws.String(filename),
		})
	if err != nil {
		logrus.Error((err))
	}
	return nil
}
