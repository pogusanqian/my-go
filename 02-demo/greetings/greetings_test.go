package greetings

import (
    "testing"
    "regexp"
)

func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
			t.Fatalf("运行错苏")
    }
}

func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("sss")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}