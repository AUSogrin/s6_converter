package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Converter(received string) (string, error) {
	received = strings.TrimSpace(received)
	if received == "" {
		return "", errors.New("empty input")
	}

	if isMorseCode(received) {
		textResult := morse.ToText(received)
		if !strings.Contains(textResult, "No encoding for") {
			return textResult, nil
		}
	}
	morseResult := morse.ToMorse(strings.ToUpper(received))
	if morseResult == "" || strings.Contains(morseResult, "No encoding for") {
		return "", errors.New("invalid input")
	}

	return morseResult, nil
}

func isMorseCode(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '/' {
			return false
		}
	}
	return true
}
