export GOOS=linux
export GOARCH=amd64
go build -o /tmp/main books
zip -j /tmp/main.zip /tmp/main
