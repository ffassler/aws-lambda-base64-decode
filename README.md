# AWS Lambda Function Base64 Decode

# Build

```shell
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

# Deploy

Deploy the zip file on AWS : main.zip  
Set the filename of the executable file : main
