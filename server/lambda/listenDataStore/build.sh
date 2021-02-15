rm -rf dist
mkdir dist
env GOOS=linux go build -o main
zip listenDataStore.zip main
mv listenDataStore.zip ./dist/
rm main
