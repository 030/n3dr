package cli

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/objectstorage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ociBackup(Bucketname string, Filename string) error {
	o, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Error("Error:", err)
		return err
	}

	ctx := context.Background()
	namespace, err := getNamespace(ctx, o)

	if err != nil {
		log.Error("Error:", err)
		return err
	}

	log.Debug("You are going to upload file " + Filename + " in bucket: " + Bucketname + "\n")
	filename, err := filepath.Glob(Filename)
	if filename == nil {
		return fmt.Errorf("error: no files found to upload")
	}

	if err != nil {
		log.Error("Error:", err)
		return err
	}
	for _, f := range filename {
		file, err := os.Open(f)
		if err != nil {
			log.Error("Error:", err)
			return err
		}
		defer file.Close()
		fi, err := file.Stat()
		if err != nil {
			log.Error("Error:", err)
			return err
		}

		err = putObject(ctx, o, namespace, Bucketname, f, fi.Size(), file, nil)
		if err != nil {
			log.Error("Error:", err)
			return err
		}

		// Removing temporary file
		if viper.GetBool("removeLocalFile") {
			err = os.Remove(f)
		}

		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func getNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) (string, error) {
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(ctx, request)
	if err != nil {
		log.Error("Error:", err)
		return "Error", err
	}
	return *r.Value, nil
}

func putObject(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, bucketname, objectname string, contentLen int64, content io.ReadCloser, metadata map[string]string) error {
	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
		ContentLength: &contentLen,
		PutObjectBody: content,
		OpcMeta:       metadata,
	}
	_, err := c.PutObject(ctx, request)
	if err != nil {
		log.Error("Error:", err)
		return err
	}
	log.Debug("You have uploaded file " + objectname + " in bucket " + bucketname + "\n")
	return err
}

func findObject(bucketname, objectname string, md5sum string) (bool, error) {
	o, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		log.Error("Error:", err)
		return false, err
	}

	ctx := context.Background()
	namespace, err := getNamespace(ctx, o)

	if err != nil {
		log.Error("Error:", err)
		return false, err
	}

	request := objectstorage.GetObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
	}
	out, err := o.GetObject(ctx, request)

	if err != nil {
		if strings.Contains(err.Error(), "ObjectNotFound") {
			return false, nil
		}
		log.Error("Error:", err)
		return false, err
	}

	md5FromOCIEncoded, _ := base64.StdEncoding.DecodeString(*out.ContentMd5)
	md5FromOCI := hex.EncodeToString(md5FromOCIEncoded)

	if md5FromOCI == md5sum {
		return true, nil
	}
	log.Debug("md5sum in OCI and md5sum on Nexus doesn't match: object will be re-created")
	return false, nil
}
