package twofa

import (
	"strings"
	"time"
)

func Generate(input string) Interface {
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	key, err := decodeKey(inputNoSpacesUpper)
	if err != nil {
		return Token{err: err}
	}
	now := time.Now()
	code := totp(key, now, 6)
	remaining := 30 - (now.Unix() % 30)
	return Token{
		code:    int(code),
		timeout: int(remaining),
	}
}
