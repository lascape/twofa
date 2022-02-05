package main

import (
	"encoding/base32"
	"fmt"
	"strings"
	"time"
)

func Generate(input string) Interface {
	inputNoSpaces := strings.Replace(input, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(inputNoSpacesUpper)
	if err != nil {
		return Token{err: err}
	}
	t := time.Now().Unix()
	code := oneTimePassword(key, toBytes(t/30))
	remaining := 30 - (t % 30)
	return Token{
		code:    int(code),
		timeout: int(remaining),
	}
}

func main() {
	generate := Generate("WRQLGCZHKY6HMEL4")
	fmt.Println(generate)
}
