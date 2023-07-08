package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"mime/multipart"
	"s3-microservice/common"
	"s3-microservice/repositories"
)

const (
	AWS_ACCESS_KEY_ID     = ""                                     // AWS access key ID
	AWS_SECRET_ACCESS_KEY = ""                                     // AWS secret access key
	AWS_BUCKET_NAME       = "65d8fe53-e3f3-463f-84b0-33858bac251f" // Default AWS bucket name
	AWS_REGION            = "us-east-2"                            // Default AWS region
)

type S3Service struct {
	userRepo repositories.UserRepository
}

func (us *S3Service) UploadFile(fileName string, file multipart.File) (string, error) {
	// S3 client configuration
	sess, _ := createAWSSession()
	svc := s3.New(sess)

	// Upload the file to S3
	result, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(common.GetEnv("AWS_BUCKET_NAME", AWS_BUCKET_NAME)),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	// Return the object version ID
	return *result.ETag, nil
}

func (us *S3Service) DownloadFile(key string) ([]byte, error) {

	// S3 client configuration
	sess, _ := createAWSSession()
	svc := s3.New(sess)

	// Download the file from S3
	output, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(common.GetEnv("AWS_BUCKET_NAME", AWS_BUCKET_NAME)),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	// Read the bytes from the S3 object
	fileBytes, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func createAWSSession() (*session.Session, error) {
	accessKey := common.GetEnv("AWS_ACCESS_KEY_ID", AWS_ACCESS_KEY_ID)
	secretKey := common.GetEnv("AWS_SECRET_ACCESS_KEY", AWS_SECRET_ACCESS_KEY)

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(common.GetEnv("AWS_REGION", AWS_REGION)),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}
