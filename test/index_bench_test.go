package index

import (
	"strings"
	"testing"
)

func BenchmarkIndexTest(b *testing.B) {
	const s = "some_text=somenvalue"
	for i := 0; i < b.N; i++ {
		strings.Index(s, "v")
	}
}

