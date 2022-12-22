package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/duck8823/sample-lambda/generated/go/user"
	"log"
)

func HandleRequest(ctx context.Context, user user.User) (string, error) {
	log.Println(fmt.Sprintf("%#v", user))
	return fmt.Sprintf("Hello %s!", user.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
