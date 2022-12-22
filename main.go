package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/duck8823/sample-lambda/generated/go/user"
	"log"
)

//go:generate openapi-generator generate -g go -i openapi.yml -o generated/go/user --package-name user --global-property=models,apis,supportingFiles=configuration.go:client.go:utils.go,apiTests=false,apiDocs=false,modelTests=false,modelDocs=false
func HandleRequest(ctx context.Context, user user.User) (string, error) {
	log.Printf("%#v\n", user)
	return fmt.Sprintf("Hello %s!", user.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
