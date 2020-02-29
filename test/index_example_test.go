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


