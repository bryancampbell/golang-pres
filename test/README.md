# Testing with GoLang (aka GO)

## Features

GoLang has many built in frameworks for making sure your code is bug-free.

* Testing
* Benchmarking
* Document Examples
* Code checks for common programmer mistakes
* CPU Profiling
* Memory Profiling

## Testing

Create a new file at ../test/index_test.go

```
package index

import (
    "strings"
    "testing"
)

func TestIndex(t *testing.T) {
    var tests = []struct {
        s   string
        sep string
        out int
    }{
        {"", "", 0},
        {"", "a", -1},
        {"fo", "foo", -1},
        {"foo", "foo", 0},
        {"oofofoofooo", "f", 2},
        // etc
    }
    for _, test := range tests {
        actual := strings.Index(test.s, test.sep)
        if actual != test.out {
            t.Errorf("Index(%q,%q) = %v; want %v", test.s, test.sep, actual, test.out)
        }
    }
}
```

The go tool runs tests.

```
$ go test
PASS
```

```
$ go test -v
=== RUN TestIndex
--- PASS: TestIndex (0.00 seconds)
PASS
```

To run the tests for all my projects:
```
$ go test github.com/bryancampbell/...
```

## Benchmarking

Create a new file at ../test/index_bench_test.go

```
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
```

The benchmark package will vary b.N until the benchmark function lasts long enough to be timed reliably.

```
$ go test -test.bench=Index
goos: darwin
goarch: amd64
BenchmarkIndexTest-8   	215134378	         5.59 ns/op
PASS
```

## Document Examples

Create a new file at ../test/index_example_test.go

```
package index

import (
    "fmt"
    "strings"
)

func ExampleIndex() {
    fmt.Println(strings.Index("chicken", "ken"))
    fmt.Println(strings.Index("chicken", "dmr"))
    // Output:
    // 4
    // -1
}
```

Examples and built and run as part of the normal test suite:

```
$ go test -v
=== RUN: ExampleIndex
--- PASS: ExampleIndex (0.00 seconds)
PASS
```

The example is displayed in godoc alongside the thing it demonstrates:

[golang.org/pkg/strings/#Index](golang.org/pkg/strings/#Index)

The benchmark package will vary b.N until the benchmark function lasts long enough to be timed reliably.

