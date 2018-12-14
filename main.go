package main

import (
	"encoding/base64"
	"github.com/aws/aws-lambda-go/lambda"
)

func Decode64(s string) (string, error) {
	value, err := base64.StdEncoding.DecodeString(s)
	return string(value), err
}

func main() {
	lambda.Start(Decode64)
}
