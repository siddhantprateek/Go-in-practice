package main

import "fmt"

// global
var g int = 20
func main(){

    // local
    var g int = 10
    fmt.Printf("%d", g)
}
