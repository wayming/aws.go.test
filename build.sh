export GOOS=linux
export GOARCH=amd64
#go build -o /tmp/main aws.go.test
#zip -j /tmp/main.zip /tmp/main


go build aws.go.test/internal/pkg/awsctl
go install aws.go.test/internal/app/awsctl