package utils

import "log"

// 致命的なエラーをログに出力する
func FatalIfErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", err)
	}
}
