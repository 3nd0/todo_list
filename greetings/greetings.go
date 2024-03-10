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

// 名前を指定された各人を関連付けるマップを返す
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) // nameとmessageを関連付けるためのmap

	// 受け取ったnameのスライスをループし、それぞれのnameのmessageを取得する。
	// Hello関数を呼び出して、それぞれのnameのmessageを取得する
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
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
