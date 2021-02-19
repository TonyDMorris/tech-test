rm -rf dist
mkdir dist
env GOOS=linux go build main.go
zip listenQuestionnaireCapture.zip main
mv listenQuestionnaireCapture.zip ./dist/
rm main
