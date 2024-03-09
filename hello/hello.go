package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	// greetingsメッセージをもらい、それを出力
	message := greetings.Hello("Glabys")
	fmt.Println(message)
}
