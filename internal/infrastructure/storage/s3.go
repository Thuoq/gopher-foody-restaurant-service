package storage

import (
	"context"
	"fmt"
	"gopher-restaurant-service/internal/config"
	"gopher-restaurant-service/internal/core/ports"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3StorageService struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
	region        string
}

func NewS3StorageService(cfg *config.Config) (ports.StorageService, error) {
	// 1. Setup credentials and config
	creds := credentials.NewStaticCredentialsProvider(cfg.S3.AccessKeyID, cfg.S3.SecretAccessKey, "")
	
	sdkConfig, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(cfg.S3.Region),
		awsConfig.WithCredentialsProvider(creds),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(sdkConfig)
	presignClient := s3.NewPresignClient(client)

	return &s3StorageService{
		client:        client,
		presignClient: presignClient,
		bucketName:    cfg.S3.BucketName,
		region:        cfg.S3.Region,
	}, nil
}

func (s *s3StorageService) GeneratePresignedURL(ctx context.Context, fileName string, contentType string) (*ports.GeneratePresignedURLOutput, error) {
	// Generate a unique key to avoid collisions
	key := fmt.Sprintf("uploads/%d-%s", time.Now().Unix(), fileName)

	presignedReq, err := s.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}, s3.WithPresignExpires(time.Minute*15))

	if err != nil {
		return nil, err
	}

	finalURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucketName, s.region, key)

	return &ports.GeneratePresignedURLOutput{
		UploadURL: presignedReq.URL,
		FinalURL:  finalURL,
	}, nil
}
