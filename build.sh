export GOOS=linux
export GOARCH=amd64

go get github.com/aws/aws-lambda-go/lambda
go get github.com/aws/aws-sdk-go
go get github.com/aws/aws-lambda-go/events

go build aws.go.test/internal/pkg/awsctl
go build aws.go.test/internal/pkg/books
go install aws.go.test/internal/app/awsctl
go install aws.go.test/internal/pkg/books

go build -o /tmp/main aws.go.test/cmds/hello
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name myfunc1 --zip-file fileb:///tmp/main.zip
aws lambda invoke --function-name myfunc1 /tmp/output.json
curl https://6xysn64cl8.execute-api.ap-southeast-2.amazonaws.com/staging/books?isbn=000-0000000000