package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b   float64
		want   float64
		hasErr bool
	}{
		{4, 2, 2, false},
		{10, 5, 2, false},
		{1, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f/%f", tt.a, tt.b), func(t *testing.T) {
			result, err := divide(tt.a, tt.b)
			if (err != nil) != tt.hasErr {
				t.Errorf("expected error: %v, error obtainedÑ %v", tt.hasErr, err)
			}
			if result != tt.want {
				t.Errorf("expected: %f, obtained: %f", tt.want, result)
			}

		})
	}
}

//for execute in console "go test"
