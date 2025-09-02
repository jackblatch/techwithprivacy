package main

import (
	"context"
	"fmt"
	"mime"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"log"
	"os"
	"techwithprivacy/web/routes"
)

const fileName = "index.html"

func getMimeType(fileName string) string {
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)

	if mimeType != "" {
		return mimeType
	}

	return "application/octet-stream"
}

func deploy(fileName string) {
	var bucketName = os.Getenv("CLOUDFLARE_R2_BUCKET_NAME")
	var accountId = os.Getenv("CLOUDFLARE_R2_ACCOUNT_ID")
	var accessKeyId = os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	var accessKeySecret = os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_SECRET")

	envVars := map[string]string{
		"CLOUDFLARE_R2_BUCKET_NAME":       bucketName,
		"CLOUDFLARE_R2_ACCOUNT_ID":        accountId,
		"CLOUDFLARE_R2_ACCESS_KEY_ID":     accessKeyId,
		"CLOUDFLARE_R2_ACCESS_KEY_SECRET": accessKeySecret,
	}

	for key, value := range envVars {
		if len(value) < 1 {
			log.Fatalf("Environment variable %s is missing", key)
		}
	}

	var bucketKey = "home.html"
	var filePath = "./" + fileName

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
		o.UseAccelerate = false
	})

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", filePath, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file %s: %v", filePath, err)
		}
	}()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &bucketKey,
		Body:        file,
		ContentType: aws.String(getMimeType(fileName)),
	})
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	fmt.Printf("File %s uploaded successfully to bucket %s with key %s\n", filePath, bucketName, bucketKey)
}

func main() {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	page, err := routes.GetIndex()

	if err != nil {
		log.Fatalf("failed to get index page: %v", err)
	}

	err = page.Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

	deploy(fileName)

}
