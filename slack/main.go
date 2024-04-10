package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slack-notify/pkg/awsapi"
	"slack-notify/pkg/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type IssueComment struct {
	URL      string `json:"url"`
	IssueURL string `json:"issue_url"`
	Body     string `json:"body"`
	User     User   `json:"user"`
}

type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	NodeID    string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
}

var (
	s3Client awsapi.S3Iface
)

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}
	s3Client = awsapi.NewS3Client(s3.NewFromConfig(sdkConfig))
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal([]byte(request.Body), &tempMap); err != nil {
		return errorResponse(err, 500)
	}

	fileName := util.GenerateFileName("comment")
	uploadPath := fmt.Sprintf("/tmp/%s", fileName)

	commentData := tempMap["comment"]

	if err := util.Create(uploadPath, commentData); err != nil {
		return errorResponse(err, 500)
	}

	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		bucketName = "sam-s3input-hisosi1900day00000" // Default bucket name
	}

	if err := s3Client.UploadFile(bucketName, "comment/"+fileName, uploadPath); err != nil {
		return errorResponse(err, 500)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func errorResponse(err error, statusCode int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       err.Error(),
		StatusCode: statusCode,
	}, nil
}
