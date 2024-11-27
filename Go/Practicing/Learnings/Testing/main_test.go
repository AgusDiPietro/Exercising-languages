package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := divide(4, 2)
	if err != nil {
		t.Errorf("Error:", err)
		return
	}
	if result != 2 {
		t.Errorf("Result different to 2: %f", result)
	}
}

//for execute in console "go test"
