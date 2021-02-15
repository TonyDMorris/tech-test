rm -rf dist
mkdir dist
env GOOS=linux go build main.go
zip listenCapture.zip main
mv listenCapture.zip ./dist/
rm main
