package greetings

import (
	"fmt"
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
func TestAllRandomFormat(t *testing.T) {
	name := "hogeo"
	formats := []string{
		"Hi, %v. Welcome (^Д^) // !!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Bonjour, %v vous allez bien ?",
	}
	counts := make(map[string]int)
	for _, format := range formats {
		counts[format] = 0
	}

	for i := 0; i < 100; i++ {
		msg, _ := Hello(name)
		for _, format := range formats {
			expected := fmt.Sprintf(format, name)
			if msg == expected {
				counts[format] += 1
			}
		}

	}
	for _, format := range formats {
		if counts[format] == 0 {
			t.Fatalf("format: \"%s\" doesn't apear in 100 tries counts: %v", format, counts)
		}
	}

}
func TestHelloNameEmpty(t *testing.T) {
	name := ""

	_, error := Hello(name)
	if error == nil {
		t.Fatal(`name is empty, error must be not null`)
	}
}
