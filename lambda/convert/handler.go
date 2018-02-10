package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

func handle(ctx context.Context, event Event) (string, error) {
	return fmt.Sprintf("Hello ctx %+v, %s!", ctx, event.Name), nil
}

func main() {
	lambda.Start(handle)
}
