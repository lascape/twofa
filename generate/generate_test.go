package generate

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	authenticator, err := Generate("WRQLGCZHKY6HMEL4")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(authenticator)
}
