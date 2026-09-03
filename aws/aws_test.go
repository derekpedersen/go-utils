package aws_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	utilsaws "github.com/derekpedersen/go-utils/aws"
)

type fakeS3 struct {
	s3iface.S3API
	body io.ReadCloser
	err  error
}

func (client fakeS3) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	if client.err != nil {
		return nil, client.err
	}
	return &s3.GetObjectOutput{Body: client.body}, nil
}

func TestDownloadWithClient(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "download.txt")
	err := utilsaws.DownloadWithClient(filename, "bucket", fakeS3{body: io.NopCloser(bytes.NewBufferString("payload"))})
	if err != nil {
		t.Fatal(err)
	}
	contents, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "payload" {
		t.Fatalf("got %q", contents)
	}
}

func TestDownloadWithClientRejectsNilClient(t *testing.T) {
	err := utilsaws.DownloadWithClient(filepath.Join(t.TempDir(), "download.txt"), "bucket", nil)
	if err == nil {
		t.Fatal("expected nil client error")
	}
}
