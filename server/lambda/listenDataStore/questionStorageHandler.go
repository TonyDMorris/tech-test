/**
A storage handler that creates a consolidated json document for each question where the
answers are batched by question over a month time period.

Here is an example of the output:

{
	"questionId": "2",
	"quizId": "0",
	"answerData": [
		{
			"answerId": "8",
			"responseFile": [
				"0/3-2020/04a60537-302e-4311-80b1-f977fede663d.json"
			]
		},
		{
			"answerId": "9",
			"responseFile": [
				"0/3-2020/04a60537-302e-4311-80b1-f977fede663d.json"
			]
		},
		{
			"answerId": "10",
			"responseFile": [
				"0/3-2020/04a60537-302e-4311-80b1-f977fede663d.json"
			]
		},
		{
			"answerId": "11",
			"responseFile": [
				"0/3-2020/04a60537-302e-4311-80b1-f977fede663d.json"
			]
		},
		{
			"answerId": "12",
			"responseFile": [
				"0/3-2020/04a60537-302e-4311-80b1-f977fede663d.json"
			]
		},
		{
			"answerId": "13",
			"responseFile": [
				"0/3-2020/04a60537-302e-4311-80b1-f977fede663d.json"
			]
		}
	]
}

This json document is stored in s3 in a folder /<questionId>/month-year/<questionId>

This means that a new file will be created every month if this question is asked!
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// QuestionData is a struct that we use to output the data for each question
type QuestionData struct {
	QuestionID string       `json:"questionId"`
	QuizID     string       `json:"quizId"`
	AnswerData []AnswerData `json:"answerData"`
}

// AnswerData is a struct that we use to store the answers within QuestionData
type AnswerData struct {
	AnswerID         string   `json:"answerId"`
	ResponseDataFile []string `json:"responseFile"`
}

// QuestionStorageHandler is a struct that implements the DataStorageHandler Interface
type QuestionStorageHandler struct {
	QuizObject QuizObject
}

// NewQuestionStorageHandler is a constructor for the QuestionStorageHandler struct
func NewQuestionStorageHandler(quizData QuizData) *QuestionStorageHandler {
	questionHandler := new(QuestionStorageHandler)
	questionHandler.QuizObject = quizData.QuizData
	return questionHandler
}

func (data QuestionStorageHandler) processData(dataFile string) {
	// Get the bucket name where we will store the processed data
	bucket := os.Getenv("DataStorage")

	for _, answers := range data.QuizObject.QuestionAnswers {

		fileName := getQuestionFileName(answers)

		// Try to load an existing file.
		questionFileBytes := loadData(bucket, fileName)

		if questionFileBytes == nil {
			data.createNewQuestionDataFile(answers, fileName, dataFile)
		} else {
			data.updateQuestionDataFile(questionFileBytes, answers, fileName, dataFile)
		}
	}
}

func getQuestionFileName(questionAnswer QuestionAnswers) string {
	year, month, _ := time.Now().Date()
	return fmt.Sprintf("%s/%d-%d/%s.json", questionAnswer.QuestionID, int(month), year, questionAnswer.QuestionID)
}

func (data QuestionStorageHandler) createNewQuestionDataFile(answers QuestionAnswers, fileName string, dataFile string) {
	questionData := new(QuestionData)
	questionData.QuizID = data.QuizObject.QuizID
	questionData.QuestionID = answers.QuestionID

	questionData.AnswerData = make([]AnswerData, len(answers.Answers))

	// This is the correct implementation
	for i, answer := range answers.Answers {
		answerData := new(AnswerData)
		answerData.AnswerID = answer
		answerData.ResponseDataFile = []string{dataFile}
		questionData.AnswerData[i] = *answerData
	}

	// This is the version in the programming challenge where we append and hence store null data
	// for _, answer := range answers.Answers {
	// 	answerData := new(AnswerData)
	// 	answerData.AnswerID = answer
	// 	answerData.ResponseDataFile = []string{dataFile}

	// 	questionData.AnswerData = append(questionData.AnswerData, *answerData)
	// }

	jsonData, err := json.Marshal(*questionData)

	if err == nil {
		storeS3(string(jsonData), fileName)
	} else {
		exitErrorf("Unable to marshal json, %v", err)
	}
}

func (data QuestionStorageHandler) updateQuestionDataFile(existingQuestionDataBytes []byte, answers QuestionAnswers, fileName string, dataFile string) {
	// Marshal Data
	existingQuestionData := parseExistingQuestionData(existingQuestionDataBytes)

	for _, answer := range answers.Answers {
		answerData := new(AnswerData)
		answerData.AnswerID = answer
		answerData.ResponseDataFile = []string{dataFile}
		existingQuestionData.AnswerData = append(existingQuestionData.AnswerData, *answerData)
	}

	jsonData, err := json.Marshal(existingQuestionData)

	if err == nil {
		storeS3(string(jsonData), fileName)
	} else {
		exitErrorf("Unable to martial json, %v", err)
	}

	fmt.Println("finished updateQuestionDataFile")
}

// ParseExistingQuestionData is a function that takes in a byte representation of QuestionData ( data formatted to analyse the responses to a question) and returns a QuestionData object
func parseExistingQuestionData(data []byte) QuestionData {
	var result QuestionData
	json.Unmarshal(data, &result)
	return result
}
