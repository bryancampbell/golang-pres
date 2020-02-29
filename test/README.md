# Testing with GoLang (aka GO)

## Features

GoLang has many built in frameworks for making sure your code is bugfree.

* Testing
* Benchmarking
* Document Examples
* Code checks for common programmer mistakes
* CPU Profiling
* Memory Profiling

## Testing

```
package index_test

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
