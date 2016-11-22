/*
* @Author: GuoDi
* @Date:   2016-04-11 22:13:41
* @Last Modified by:   GuoDi
* @Last Modified time: 2016-11-11 01:55:40
 */
package aliyun

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

const (
	endpoint        = beego.AppConfig.String("aliyun_endpoint")
	AccessKeyID     = beego.AppConfig.String("aliyun_access_key_id")
	AccessKeySecret = beego.AppConfig.String("aliyun_access_key_secret")
	bucketName      = beego.AppConfig.String("aliyun_oss_bucket_name")
)

var (
	client *oss.Client
)

func init() {
	var err error
	client, err = oss.New(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("oss warning:", err.Error())
	}
}
func GetBucket(name string) *oss.Bucket {
	if name == "" {
		name = bucketName
	}
	bucket, err := client.Bucket(name)
	if err != nil {
		fmt.Println("bucket warning:", err.Error())
		return nil

	}
	return bucket
}
func PutFile(name string, file string) bool {
	if bucket := GetBucket(""); bucket != nil {
		err := bucket.PutObjectFromFile(name, file)
		if err != nil {
			fmt.Println("file put warning:", err.Error())
			return false

		}
		return true
	}
	return false
}
func PutObject(name string, obj []byte) bool {
	if bucket := GetBucket(""); bucket != nil {
		err := bucket.PutObject(name, bytes.NewReader(obj))
		if err != nil {
			fmt.Println("file put warning:", err.Error())
			return false

		}
		return true
	}
	return false
}
func PutObjectFile(name string, fd multipart.File) bool {
	if bucket := GetBucket(""); bucket != nil {
		err := bucket.PutObject(name, fd)
		if err != nil {
			fmt.Println("file put warning:", err.Error())
			return false

		}
		return true
	}
	return false
}
