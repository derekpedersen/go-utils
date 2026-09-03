package aws

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func Download(filename, s3bucket string) error {
	if AWS_SESSION == nil {
		return os.ErrInvalid
	}
	return DownloadWithClient(filename, s3bucket, s3.New(AWS_SESSION))
}

func DownloadWithClient(filename, s3bucket string, client s3iface.S3API) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := file.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	if client == nil {
		return os.ErrInvalid
	}
	object, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return err
	}
	defer object.Body.Close()
	_, err = io.Copy(file, object.Body)
	return err
}
