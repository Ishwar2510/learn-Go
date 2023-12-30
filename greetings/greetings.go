package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	formats := []string{
		"Hey %v",
		"I miss u , %v",
		"I love you, %v",
	}

	// message := fmt.Sprintf("Hi %v, Welcome", name)
	// return message, nil

	return fmt.Sprintf(formats[rand.Intn(len(formats))], name), nil
}
