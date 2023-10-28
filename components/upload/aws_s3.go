package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(
	accessKey string,
	secretKey string,
	bucketName string,
	region string,
	domain string,
) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     accessKey,
		secret:     secretKey,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secret, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider
}

func (provider *s3Provider) Upload(ctx context.Context, data []byte, folder, filename string) (string, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	dst := folder + "/" + filename

	fmt.Println("fileType====>", fileType)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s/%s", provider.domain, dst)

	return url, err
}
