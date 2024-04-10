package util

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func Create(fileName string, data json.RawMessage) error {
	// ファイルを作成します。ファイルがすでに存在する場合は、空のファイルにトランケートされます。
	file, err := os.Create(fileName)
	if err != nil {
		// エラー処理
		log.Fatal(err)
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GenerateFileName(fileName string) string {
	// 現在の時間を取得
	now := time.Now()
	// Unixタイムスタンプを取得し、それを基にランダムな値を生成
	src := rand.NewSource(now.UnixNano())
	r := rand.New(src)

	// 0から9999の範囲でランダムな数を生成
	randomPart := r.Intn(10000)

	// 時間をYYYYMMDDHHMMSS形式の文字列に変換
	timeStr := now.Format("20060102150405")

	// ファイル名を生成
	return fmt.Sprintf("%s-%s-%d", fileName, timeStr, randomPart)
}
