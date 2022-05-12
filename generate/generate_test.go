package generate

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	secret := Secret()
	fmt.Println(secret, len(secret))

	authenticator := Generate(secret)
	t.Log(authenticator)

	alignment := codeAlignment(12345)
	t.Log(alignment)
}
