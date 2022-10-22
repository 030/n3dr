package artifacts

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/objectstorage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ociObject struct {
	ctx        context.Context
	c          objectstorage.ObjectStorageClient
	namespace  string
	bucketname string
	objectname string
	contentLen int64
	content    io.ReadCloser
	metadata   map[string]string
}

func ociBackup(Bucketname, Filename string) (errs []error) {
	o, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	ctx := context.Background()
	namespace, err := getNamespace(ctx, o)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	log.Debug("You are going to upload file " + Filename + " in bucket: " + Bucketname + "\n")
	filename, err := filepath.Glob(Filename)
	if filename == nil {
		err = fmt.Errorf("error: no files found to upload")
		errs = append(errs, err)
	}

	if err != nil {
		errs = append(errs, err)
		return errs
	}
	for _, f := range filename {
		file, err := os.Open(filepath.Clean(f))
		if err != nil {
			errs = append(errs, err)
			return errs
		}
		defer func() {
			if err := file.Close(); err != nil {
				errs = append(errs, err)
			}
		}()

		fi, err := file.Stat()
		if err != nil {
			errs = append(errs, err)
			return errs
		}

		o := ociObject{ctx, o, namespace, Bucketname, f, fi.Size(), file, nil}
		err = o.putObject()
		if err != nil {
			errs = append(errs, err)
			return errs
		}

		// Removing temporary file
		if viper.GetBool("removeLocalFile") {
			err = os.Remove(f)
		}

		if err != nil {
			errs = append(errs, err)
			return errs
		}
	}
	return nil
}

func getNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) (string, error) {
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(ctx, request)
	if err != nil {
		return "Error", err
	}
	return *r.Value, nil
}

func (o ociObject) putObject() error {
	request := objectstorage.PutObjectRequest{
		NamespaceName: &o.namespace,
		BucketName:    &o.bucketname,
		ObjectName:    &o.objectname,
		ContentLength: &o.contentLen,
		PutObjectBody: o.content,
		OpcMeta:       o.metadata,
	}
	_, err := o.c.PutObject(o.ctx, request)
	if err != nil {
		return err
	}
	log.Debug("You have uploaded file " + o.objectname + " in bucket " + o.bucketname + "\n")
	return err
}

func findObject(bucketname, objectname, md5sum string) (bool, error) {
	o, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		return false, err
	}

	ctx := context.Background()
	namespace, err := getNamespace(ctx, o)

	if err != nil {
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
