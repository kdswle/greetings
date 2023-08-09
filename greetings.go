package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// マップ内で、
		// 取得したメッセージを名前に関連付けます。
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome (^Д^) // !!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Bonjour, %v vous allez bien ?",
	}
	return formats[rand.Intn(len(formats))]
}
