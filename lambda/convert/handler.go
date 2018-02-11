package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mikeflynn/go-alexa/skillserver"
)

func handle(ctx context.Context, er *skillserver.EchoRequest) (*skillserver.EchoResponse, error) {
	log.Printf("Hello requeset %+v", er)
	resp := skillserver.NewEchoResponse()
	resp.Card("test", "test")
	return resp, nil
}

func main() {
	lambda.Start(handle)
}
