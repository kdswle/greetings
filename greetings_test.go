package greetings

import (
	"strings"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "hogeo"
	msg, _ := Hello(name)
	if !strings.Contains(msg, name) {
		t.Fatalf(`Hello("Gladys") = %q, must contains %q`, msg, name)
	}
}
func TestHelloNameEmpty(t *testing.T) {
	name := ""

	_, error := Hello(name)
	if error == nil {
		t.Fatal(`name is empty, error must be not null`)
	}
}
