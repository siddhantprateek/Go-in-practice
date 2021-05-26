package main

import "fmt"

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
