package greetings

import (
	"fmt"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "hogeo"
	want := fmt.Sprintf("Hi, %s. Welcome (^Д^) // !!", name)
	msg, _ := Hello(name)
	if msg != want {
		t.Fatalf(`Hello("Gladys") = %q, want match for %q, nil`, msg, want)
	}
}
func TestHelloNameEmpty(t *testing.T) {
	name := ""

	_, error := Hello(name)
	if error == nil {
		t.Fatal(`name is empty, error must be not null`)
	}
}
