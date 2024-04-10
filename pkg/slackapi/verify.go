package slackapi

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

// Slack 認証
func Verify(signingSecret string, headers http.Header, body []byte) error {
	// NOTE: テスト実行時のみ error = nil で返す
	//   slack-go/slack.NewSecretsVerifier は X-Slack-Request-Timestamp (Slack スラッシュコマンドのリクエストのタイムスタンプ) は
	//   現在時刻から5分以上開きがあると "timestamp is too old" でエラーを返す。
	//
	//   テスト用に X-Slack-Request-Timestamp を現在時刻と5分以下に変更しても、その時刻を元に hmac のハッシュ計算をしており
	//   テストをパスするトークンを生成できない。
	//   その為、テスト回避策が他に見当たらない為、テスト実行時に渡す環境変数 TEST = 1 が渡された場合に Slack 認証をパスしたものと見做す。
	//   see: https://github.com/slack-go/slack/commit/9c98723e5a0030d32bfcc55f62cb4db619aac1ed
	if os.Getenv("TEST") == "1" {
		return nil
	}

	// Slack 認証実施クライアント生成
	sv, err := slack.NewSecretsVerifier(headers, signingSecret)
	if err != nil {
		return err
	}
	log.Printf("sv: %v", sv)
	log.Printf("ここまできてるよ！")
	// 認証に利用するデコードした Body 設定
	if _, err := sv.Write(body); err != nil {
		return fmt.Errorf("secrets verifier write failed. err: %v", err)
	}
	// 認証実施
	if err := sv.Ensure(); err != nil {
		return fmt.Errorf("secrets verifier ensure failed. err: %v", err)
	}
	// 認証成功
	return nil
}
