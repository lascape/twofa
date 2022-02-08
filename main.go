package main

import (
	"flag"
	"fmt"

	"github.com/golandscape/twofa/generate"
)

func main() {
	secret := flag.String("secret", "", "秘钥")
	flag.Parse()

	if *secret == "" {
		fmt.Println("secret must not empty")
	}

	authenticator, err := generate.Generate(*secret)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("result:%+v\n", authenticator)
}
