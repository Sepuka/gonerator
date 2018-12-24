package number

import (
	"testing"
)

func BenchmarkUintOnlyRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SeqUint64(1e5)
	}
}

func BenchmarkUintGeneratorOnlyGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UintGenerator(1e5)
	}
}
