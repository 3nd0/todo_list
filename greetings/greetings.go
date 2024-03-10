package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// 指定された人物の挨拶を返す
func Hello(name string) (string, error) {
	// 名前が渡されなければエラーと一緒にメッセージを返す
	if name == "" {
		return "", errors.New("empty name")
	}

	// ランダムフォーマットでメッセージを生成する
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Wellcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// フォーマットのスライスのランダムなインデックスを指定して、ランダムに選択されたメッセージ・フォーマットを返す。
    // フォーマットのスライスに対するランダムなインデックスを指定する。
	return formats[rand.Intn(len(formats))]
}