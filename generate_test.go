package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	generate := Generate("XXQLGCZHKY6HMEUU")
	t.Log(generate)
}
