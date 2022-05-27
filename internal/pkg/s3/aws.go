package s3

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	log "github.com/sirupsen/logrus"
)

type Nexus3 struct {
	AwsBucket, AwsId, AwsRegion, AwsSecret, ZipFilename string
}

func (n *Nexus3) Upload() error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(n.AwsRegion), Credentials: credentials.NewStaticCredentials(n.AwsId, n.AwsSecret, "")})
	if err != nil {
		return fmt.Errorf("session.NewSession - filename: %v, err: %v", n.ZipFilename, err)
	}
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(n.ZipFilename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", n.ZipFilename, err)
	}

	filename := filepath.Base(n.ZipFilename)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(n.AwsBucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	log.Infof("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil
}
