package aws

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/globalsign/mgo/bson"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type S3Credential struct {
	AWSConfig aws.Config
	Bucket    string
	Directory string
	AccessKey string
	SecretKey string
}

//function s3 upload manager to s3 storage
func (s3Cred S3Credential) UploadManager(fileToBeUploaded *multipart.FileHeader) (res string, err error) {
	//initialization aws session
	session, err := awsSession.NewSession(&s3Cred.AWSConfig)
	if err != nil {
		return res, err
	}

	//open file
	file, err := fileToBeUploaded.Open()
	if err != nil {
		return res, err
	}

	//get file size to set content, and rename file name.for setting in s3 putObjectInput struct
	size := fileToBeUploaded.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	fileName := bson.NewObjectId().Hex() + filepath.Ext(fileToBeUploaded.Filename)
	res = s3Cred.Directory + "/" + fileName

	//put file to s3 storage
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(buffer),
		Bucket:             aws.String(s3Cred.Bucket),
		ContentDisposition: aws.String("attachment"),
		ContentLength:      aws.Int64(int64(size)),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		Key:                aws.String(res),
	})
	if err != nil {
		return res, err
	}

	return res, err
}

//function to get pre-signed url from S3, pre-signed url is encrypted url from s3 that present our file in s3
func (s3Cred *S3Credential) GetURL(key string) (res string, err error) {
	//initialization aws session
	sess, err := awsSession.NewSession(&s3Cred.AWSConfig)
	if err != nil {
		return res, err
	}

	//open aws session to get pre-signed url from url
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Cred.Bucket),
		Key:    aws.String(key),
	})

	//get pre-signed url with 15 minutes limit expiration link
	res, err = req.Presign(15 * time.Minute)

	return res, err
}

func (s3Cred S3Credential) UploadManagerBufferWithSession(session *awsSession.Session, buffer []byte, fileName string) (res string, err error) {
	//get file size to set content, and rename file name.for setting in s3 putObjectInput struct
	bucketFilePath := s3Cred.Directory + "/" + fileName

	//put file to s3 storage
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		ACL:                aws.String("public-read"),
		Body:               bytes.NewReader(buffer),
		Bucket:             aws.String(s3Cred.Bucket),
		ContentDisposition: aws.String("attachment"),
		ContentLength:      aws.Int64(int64(len(buffer))),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		Key:                aws.String(bucketFilePath),
	})
	if err != nil {
		return res, err
	}

	res = bucketFilePath
	return res, err
}

func (s3Cred *S3Credential) GetURLWithSession(session *awsSession.Session, key string) (res string, err error) {
	// create svc
	svc := s3.New(session)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Cred.Bucket),
		Key:    aws.String(key),
	})

	//get pre-signed url with 15 minutes limit expiration link
	res, err = req.Presign(15 * time.Minute)

	return res, err
}

func (s3Cred *S3Credential) IsExistWithSession(session *awsSession.Session, key string) (bool, error) {
	_, err := s3.New(session).HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(s3Cred.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if errCode, ok := err.(awserr.Error); ok {
			switch errCode.Code() {
			case "NotFound": // s3.ErrCodeNoSuchKey does not work, aws is missing this error code so we hardwire a string
				return false, nil
			default:
				return false, err
			}
		}
		return false, err
	}
	return true, nil
}

func (s3Cred *S3Credential) Download(key, path string) (res string, err error) {

	sess, err := awsSession.NewSession(&s3Cred.AWSConfig)
	if err != nil {
		return res, err
	}

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(path + `/` + key)
	if err != nil {
		return res, err
	}

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s3Cred.Bucket),
			Key:    aws.String(key),
		})

	res = file.Name()

	return res, err
}
