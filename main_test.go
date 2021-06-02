package myid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	// TODO: Add tests.
}

func BenchmarkGenerateUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Generate()
	}
}
