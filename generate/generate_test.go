package generate

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	authenticator := Generate("WRQLGCZHKY6HMEL4")
	t.Log(authenticator)

	alignment := codeAlignment(12345)
	t.Log(alignment)
}
