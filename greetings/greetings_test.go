package greetings

import (
	"testing"
	"regexp"
)

// TestHelloNameは、名前を指定してgreetings.Helloを呼び出し、チェックする
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	message, error := Hello("Gladys")
	if !want.MatchString(message) || error != nil {
		t.Fatalf(`Hello("Gladys) = %q, %v, want match for %#q, nil`, message, error, want)
	}
}

// TestHelloEmptyは空の文字列でgreetings.Helloを呼び出す
func TestHelloEmpty(t *testing.T) {
	message, error := Hello("")
	if message != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, error)
	}
}