package generate

import (
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	//secret := Secret()
	//fmt.Println(secret, len(secret))
	secret := "GBSTCMBSGYZTENBX"
	now := time.Now()
	for i := 0; i < 3; i++ {
		authenticator := Generate(secret, now)
		t.Log(authenticator)
		now = now.Add(-time.Second * 30)
	}

	alignment := codeAlignment(12345)
	t.Log(alignment)
}
