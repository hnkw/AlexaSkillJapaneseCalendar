package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mikeflynn/go-alexa/skillserver"
	"github.com/ushios/gengo"
)

func convert(t time.Time) (nengo string, no int, err error) {
	defer func() {
		recoverErr := recover()
		if recoverErr != nil {
			err = errors.New("panic orrurs")
			return
		}
	}()
	g := gengo.At(t)
	nengo = g.String()
	no = t.Year() - g.StartAt().Year() + 1
	return
}

func parseTime(amzdate string) (time.Time, error) {
	amzdate = strings.Replace(amzdate, "XX", "01", 2)
	return time.Parse("2006-01-02", amzdate)
}

func outputSpeech(t time.Time, nengo string, no int) string {
	return fmt.Sprintf("西暦%d年は、%s%d年です。", t.Year(), nengo, no)
}

func outputCard(t time.Time, nengo string, no int) (string, string) {
	return fmt.Sprintf("%s%d年", nengo, no), fmt.Sprintf("西暦%d年は、%s%d年です。", t.Year(), nengo, no)
}

func handle(ctx context.Context, er *skillserver.EchoRequest) (*skillserver.EchoResponse, error) {
	amzdate, err := er.GetSlotValue("year")
	if err != nil {
		return nil, err
	}

	fmt.Printf("date, %s", amzdate)

	t, err := parseTime(amzdate)
	if err != nil {
		return nil, err
	}

	nengo, no, err := convert(t)
	if err != nil {
		return nil, err
	}

	resp := skillserver.NewEchoResponse()
	resp.OutputSpeech(outputSpeech(t, nengo, no))
	resp.Card(outputCard(t, nengo, no))
	return resp, nil
}

func main() {
	lambda.Start(handle)
}
