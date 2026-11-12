package aws

import (
	"context"
	"io"
	"os"
	"path/filepath"

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

func DownloadWithClient(filename, s3bucket string, client s3iface.S3API) error {
	return DownloadKeyWithClient(filename, filename, s3bucket, client)
}

func DownloadKey(localFilename, objectKey, s3bucket string) error {
	if AWS_SESSION == nil {
		return os.ErrInvalid
	}
	return DownloadKeyWithClient(localFilename, objectKey, s3bucket, s3.New(AWS_SESSION))
}

func DownloadKeyWithContext(ctx context.Context, localFilename, objectKey, s3bucket string, client s3iface.S3API) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	return DownloadKeyWithClient(localFilename, objectKey, s3bucket, client)
}

func DownloadKeyWithClient(localFilename, objectKey, s3bucket string, client s3iface.S3API) (err error) {
	if client == nil {
		return os.ErrInvalid
	}
	if err := os.MkdirAll(filepath.Dir(localFilename), 0755); err != nil {
		return err
	}

	file, err := os.Create(localFilename)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := file.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	object, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}
	defer object.Body.Close()
	_, err = io.Copy(file, object.Body)
	return err
}
