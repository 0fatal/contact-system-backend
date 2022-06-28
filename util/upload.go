package util

import (
	"context"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
)

// upload
// 上传附件
var client *cos.Client

func init() {
	u, _ := url.Parse("")
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "",
			SecretKey: "",
		},
	})
}

func Upload(file *multipart.FileHeader) (_url string, err error) {
	fd, err := file.Open()
	if err != nil {
		return
	}
	hashs := uuid.New().String() + path.Ext(file.Filename)
	_, err = client.Object.Put(context.Background(), hashs, fd, nil)
	if err != nil {
		return
	}
	_url = client.Object.GetObjectURL(hashs).String()
	return
}
