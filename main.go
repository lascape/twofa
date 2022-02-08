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

	authenticator, err := generate.Generate(secret)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("result:%+v\n", authenticator)
}
