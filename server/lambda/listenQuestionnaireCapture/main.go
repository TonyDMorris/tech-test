package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const format = ".json"

// QuestionType placeholder for the question type type
type QuestionType string

// all possible types of a question
const (
	QuestionTypeButton      QuestionType = "button"
	QuestionTypeMultiSelect              = "multi-select"
	QuestionTypeText                     = "text"
)

type Questionnaire struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

type Question struct {
	ID      int          `json:"id"`
	Type    QuestionType `json:"type"`
	Title   string       `json:"title "`
	Answers []Answer     `json:"answers"`
}

type Answer struct {
	ID             int     `json:"id"`
	Text           string  `json:"text"`
	NextQuestionId *string `json :"next_question_id"`
}

func route(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodGet:
		return handleGetRequest(ctx, request)
	case http.MethodPost:
		return handlePostRequest(ctx, request)
	default:
		return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed}, nil
	}

}

//HandleGetRequest returns the latest quiz to the client
func handleGetRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get latest file id
	fileID := getLatestFilename()

	// SprintF into actual filename
	filename := fmt.Sprintf("%v%v", fileID, format)
	// get file bytes from s3
	fileBytes := loadData(os.Getenv("questionnaireBucket"), filename)

	return clientResponse(string(fileBytes), http.StatusOK), nil
}

// handlePostRequest validates and stores incoming questionnaires
func handlePostRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// unmarshal the questionnaire
	var Questionnaire Questionnaire
	err := json.Unmarshal([]byte(request.Body), &Questionnaire)
	if err != nil || !Questionnaire.Valid() {
		return clientResponse("the questionnaire data was malformed", http.StatusUnprocessableEntity), nil
	}

	// generate serial filename
	fileName := time.Now().Unix()

	storeS3(request.Body, fmt.Sprintf("%v%v", fileName, format))

	return clientResponse("successfully posted questionnaire", http.StatusCreated), nil
}

func storeS3(data string, filename string) {
	bucket := os.Getenv("SurveyBucket")
	fmt.Printf("bucket: %s/n", bucket)

	sess, err := session.NewSession()
	uploader := s3manager.NewUploader(sess)

	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(filename),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: strings.NewReader(data),
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func clientResponse(msg string, code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       msg,
	}
}

func getLatestFilename() int64 {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	params := &s3.ListObjectsInput{
		Bucket: aws.String(os.Getenv("QuestionnaireBucket")),
		Prefix: aws.String("/"),
	}

	resp, err := svc.ListObjects(params)
	if err != nil {
		exitErrorf("unable to retrive filename from s3", err)
	}
	// the questionnairs are stored via unix time stamp so we get the latest file by pulling the file names within the bucket
	// and math.Max the float representations until we get the latest
	var fileTime float64

	for _, key := range resp.Contents {
		fileName := *key.Key
		fileExtension := filepath.Ext(*key.Key)
		fileName = fileName[0 : len(fileName)-len(fileExtension)]
		timeFloat, err := strconv.ParseFloat(fileName, 64)
		if err != nil {
			exitErrorf("unable to parse intiger from filename", err)
		}
		fileTime = math.Max(timeFloat, fileTime)

	}
	// we return the latest filename
	return int64(fileTime)
}

func loadData(bucketName string, key string) []byte {
	svc := s3.New(session.New())

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		fmt.Printf("Error = %+v\n", err)
		fmt.Println("Error Bucket =", bucketName)
		fmt.Println("Error Key=", key)
		return nil
	}

	s3objectBytes, err := ioutil.ReadAll(result.Body)

	return s3objectBytes
}

// Valid contains the validation logic for the questionnaire , currently  not implemented
func (q *Questionnaire) Valid() bool { return true }

func main() {
	lambda.Start(route)
}
