package main

import (
	"testing"
)

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = divide(10, 2)
	}
}

//encuentran cuellos de botella en el programa, ayudan a ver su rendimiento.
