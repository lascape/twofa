package twofa

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	generate := Generate("NTHEE6VLH7DLFRSP")
	t.Log(generate)
}
