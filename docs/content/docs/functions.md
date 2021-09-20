---
title: "Functions"
weight: 3
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: false
# bookComments: false
# bookSearchExclude: false
---

# Functions

A function is a block of statements that can be used repeatedly in a program. It not execute automatically when a page loads and will be executed by a call to the function.

## Creating a function
1. It starts with key word `func`

Basic Syntax
```go

func <function-Name> (<parameters>) <return-type> {
    // code here//
}
```
> Note: `main` function is the entry point to go program like in C and C++. The main function does have a return type and it does not accept any parameters.

```go
func main() {
    // code here //
}
```


As simple program to find the fibonacci sequence .
Inside `fibo.go` file

```go
func fibo(n int) (ret int) {
    if n < 2{
        return 1
    }
    return fibo(n - 1) + fibo(n - 2)
}

func main() {

    var i int
    for i = 0; i < 10; i++{
        fmt.Printf("%d ", fibo(i))
    }
    fmt.Println("")
    for n := 0; n < 10; n++{
        fmt.Printf("%d ", fibo(n))
    }
}

```