package handlers

import (
	"fmt"
	"os"
  "strings"	

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AWS sesssion wrapper
type AwsClient struct {
	session *session.Session
}

// CreateNewSession creates an aws s3 session
func CreateNewAwsClient() (*AwsClient, error) {
	client := new(AwsClient)

	// support non-aws (ceph, swift, minio,...) endpoints
	endpoint, ok := os.LookupEnv("AWS_ENDPOINT")
	var s3Config *aws.Config
	if ok {
		s3Config = &aws.Config{
			Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
			Endpoint:         aws.String(endpoint),
			Region:           aws.String(os.Getenv("AWS_REGION")),
			DisableSSL:       aws.Bool(strings.Contains("https", endpoint)),
			S3ForcePathStyle: aws.Bool(true),
		}
	} else {
		s3Config = &aws.Config{
			Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
			Region: aws.String(os.Getenv("AWS_REGION")),
		}
	}

	sess, err := session.NewSession(s3Config)

	if err != nil {
		return nil, err
	}

	client.session = sess
	return client, nil
}

func GetS3BucketOwner(client *AwsClient, bucket string) (string, error) {
	svc := s3.New(client.session)
	input := &s3.GetBucketAclInput{
		Bucket: aws.String(bucket),
	}

	result, err := svc.GetBucketAcl(input)
	if err != nil {
		return "", err
	}

	return *result.Owner.DisplayName, nil
}

// GetS3ObjectOutput gets object output from s3
func GetS3ObjectOutput(client *AwsClient, bucket string, key string) (*s3.GetObjectOutput, error) {

	svc := s3.New(client.session)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, nil

}

// GetObjectSize returns object size in bytes
func GetObjectSize(client *AwsClient, bucket string, key string) (*int64, error) {
	svc := s3.New(client.session)
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := svc.HeadObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil, err
	}
	return result.ContentLength, nil
}
