package main

import (
	"fmt"
	"os"

	"github.com/golandscape/twofa/generate"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("secret must not empty")
		return
	}

	secret := os.Args[1]

	authenticator := generate.Generate(secret)
	if authenticator.Error != nil {
		panic(authenticator.Error)
	}
	fmt.Printf("result:%+v\n", authenticator)
}
