package greetings

import "fmt"

// 指定された人物の挨拶を返す
func Hello(name string) string {
	// nameを埋め込んだ挨拶を返す
	message := fmt.Sprintf("Hi, %v. welcome", name)
	return message
}
