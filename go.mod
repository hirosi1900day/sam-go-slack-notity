require (
	github.com/aws/aws-lambda-go v1.36.1
	github.com/aws/aws-sdk-go v1.51.14
	github.com/aws/aws-sdk-go-v2/config v1.27.11
	github.com/aws/aws-sdk-go-v2/service/s3 v1.53.1
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/slack-go/slack v0.12.5
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/sync v0.1.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8

module slack-notify

go 1.16
