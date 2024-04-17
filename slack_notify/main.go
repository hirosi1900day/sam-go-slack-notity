package main

import (
	"fmt"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/sync/errgroup"
)

func main() {
	lambda.Start(handler)
}

// SQS イベントでトリガーされ handler 実行する為、引数に events.SQSEvent を指定している
func handler(ctx context.Context, e events.SQSEvent) error {
	var eg errgroup.Group

	// SQS レコード毎に処理を実行する
	for _, m := range e.Records {
		body := m.Body
		eg.Go(func() error {
			return execute(body)
		})
	}

	// 上記、並列処理でエラーが1つでもあれば、エラーを返す
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func execute(body string) error {
	// logに出力
	fmt.Println(body)
	return nil
}
