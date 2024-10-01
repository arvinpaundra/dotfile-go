package client

import (
	"context"
	"fmt"
	"io"
	path "path/filepath"
	"strings"
	"time"

	config "github.com/arvinpaundra/dotfile-go/config"
	"github.com/arvinpaundra/dotfile-go/pkg/util"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type CloudStorageOption struct {
	Filename string
	Filepath string
}

type CloudStorage interface {
	Upload(ctx context.Context, r io.Reader, option CloudStorageOption) (string, error)
}

func (opt CloudStorageOption) GetObjectName() string {
	ext := path.Ext(opt.Filename)
	filepath := opt.Filepath
	filename := util.RandomString(32) + ext

	if filepath != "" && filepath[:1] == "/" {
		filepath = strings.Replace(filepath, "/", "", 1)
	}

	return path.Clean(filepath + "/" + filename)
}

type gcsClient struct {
	client  *storage.Client
	bucket  string
	baseurl string
}

func NewGoogleCloudStorageClient() CloudStorage {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(config.C.GCS.CredentialPath))
	if err != nil {
		panic(err)
	}

	return &gcsClient{
		client:  client,
		bucket:  config.C.GCS.Bucket,
		baseurl: config.C.GCS.BaseUrl,
	}
}

func (gcs *gcsClient) Upload(ctx context.Context, r io.Reader, option CloudStorageOption) (string, error) {
	objectName := option.GetObjectName()

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*30)

	defer cancel()

	attrs, err := gcs.client.Bucket(gcs.bucket).Attrs(timeoutCtx)
	if err != nil {
		return "", err
	}

	wc := gcs.client.Bucket(gcs.bucket).Object(objectName).NewWriter(timeoutCtx)

	if len(attrs.ACL) > 0 {
		wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	}

	_, err = io.Copy(wc, r)
	if err != nil {
		return "", err
	}

	err = wc.Close()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", gcs.baseurl, gcs.bucket, wc.Attrs().Name)

	return url, nil
}
