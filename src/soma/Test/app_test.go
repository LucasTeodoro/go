package main

import(
	"testing"
	"soma/functions"
)

func TestSoma(t *testing.T) {
	result := functions.Soma(5, 5)
	if result != 10 {
		t.Error("Valor diferente de 10")
	}
}