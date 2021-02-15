rm -rf dist
mkdir dist
env GOOS=linux go build main.go
zip listenQuestionTrigger.zip main
mv listenQuestionTrigger.zip ./dist/
rm main