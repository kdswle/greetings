package greetings

import (
	"fmt"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "hogeo"
	want := fmt.Sprintf("Hi, %s. Welcome!", name)
	msg := Hello(name)
	if msg != want {
		t.Fatalf(`Hello("Gladys") = %q, want match for %q, nil`, msg, want)
	}
}
