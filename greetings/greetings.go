package greetings

import (
	"errors"
	"fmt"
)

// 指定された人物の挨拶を返す
func Hello(name string) (string, error) {
	// 名前が渡されなければエラーと一緒にメッセージを返す
	if name == "" {
		return "", errors.New("empty name")
	}

	// nameを受け取った場合は、nameをmessageに埋め込む
	message := fmt.Sprintf("Hi, %v. welcome", name)
	return message, nil
}
