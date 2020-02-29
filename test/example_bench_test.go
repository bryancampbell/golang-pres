package index_test

import (
	"strings"
	"testing"
)

func BenchmarkIndex(b *testing.B) {
	const s = "some_text=somenvalue"
	for i := 0; i < b.N; i++ {
		strings.Index(s, "v")
	}
}

