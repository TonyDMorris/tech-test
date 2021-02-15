#!/bin/bash
# A shell script to build all of the source files and copy them to the dist directory
rm -rf dist
mkdir dist

# Build the listenCaptureLambda - This is the lambda that captures responses to questions
echo "building listenCaptureLambda"
(cd ./server/lambda/listenCapture/ && ./build.sh)
cp ./server/lambda/listenCapture/dist/listenCapture.zip ./dist

# Build the listenDataStoreLambda - This is the Lambda function that listens to the FIFO queue and stores the data into s3
echo "building listenDataStoreLambda"
(cd ./server/lambda/listenDataStore/ && ./build.sh)
cp ./server/lambda/listenDataStore/dist/listenDataStore.zip ./dist

# Build the listenQuestionTrigger 
echo "building listenQuestionTrigger"
(cd ./server/lambda/listenQuestionTrigger/ && ./build.sh)
cp ./server/lambda/listenQuestionTrigger/dist/listenQuestionTrigger.zip ./dist


