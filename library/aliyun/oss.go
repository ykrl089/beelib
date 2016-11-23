/*
* @Author: GuoDi
* @Date:   2016-04-11 22:13:41
* @Last Modified by:   GuoDi
* @Last Modified time: 2016-11-23 23:27:03
 */
package aliyun

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

var (
	client                                             *oss.Client
	Endpoint, AccessKeyID, AccessKeySecret, BucketName string
)

func init() {
	Endpoint = beego.AppConfig.String("aliyun_endpoint")
	AccessKeyID = beego.AppConfig.String("aliyun_access_key_id")
	AccessKeySecret = beego.AppConfig.String("aliyun_access_key_secret")
	BucketName = beego.AppConfig.String("aliyun_oss_bucket_name")
	if err := SetClient(); err != nil {
		fmt.Println("oss warning:", err.Error())
	}
}
func SetClient() error {
	if client, err = oss.New(Endpoint, AccessKeyID, AccessKeySecret); err != nil {
		return err
	} else {
		return nil
	}
}
func GetBucket(name string) *oss.Bucket {
	if name == "" {
		name = BucketName
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
