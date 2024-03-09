package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// 以下を含む、定義済みロガーのプロパティを設定する
	// ログ・エントリーのprefixと出力を無効にするフラグ
	// 時間、ソースファイル、行番号
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// greetingsメッセージをもらい、それを出力
	message, error := greetings.Hello("")

	// エラーが返された場合は、それをコンソールに表示しプログラムを終了する。
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
