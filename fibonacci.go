package main

import (
    "fmt"
)

func main() {
    f := fib()
    for i := 0; i < 50; i++ {
        fmt.Println(f())
    }
}

// fibonacci func
func fib() func() int64 {
    var a, b int64 = 0, 1
    return func() int64 {
        a, b = b, a+b
        return a
    }
}

