# Working with GoLang (aka GO)

## Setting up GoLang

There are a few ways to install GoLang (pick one)

- ***On Mac***

```
brew install go
```

- ***On Windows***

 1. Download Go from here https://golang.org/dl/
 2. Install Go using the installer


- ***Web Based Compiler (no need to install anything)***

 1. Open a browser
 2. Go here https://play.golang.org/


## How Do I Start Writing Code

You need to setup your GOPATH, which is where all of your code will reside

- Mac:

```
export GOPATH=/Users/mack/go

```

### Let's Write a Simple Program

Open up your favorite editor and create a new file, put the content list below in the file and save it in /Users/mack/go/src/example/main.go

```
package main

import "fmt"

func main() {
    fmt.Println("Hello Detroit Black Tech")
}
```

### Execute Your First Go Program

```
cd /Users/mack/go/src/example
go run main.go
```


## Value of GoLang - #1 - Feels likes a scripting language

Examples of GoLang

```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

A function can return any number of results. The swap function returns two strings.

```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

## Value of GoLang - #2 - Thread and Concurrency Built In


```
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
(The data flows in the direction of the arrow.)
```

Like maps and slices, channels must be created before use:

```
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

### Exercise 

```
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
  
	for {
		select {
    // Write cases for the bomb to go off.
		}
	}
}

```

The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:

```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

## Value of GoLang - #3 - Testing, Benchmarking, Profiling

[Documentation](../../blob/master/test/README.md)
