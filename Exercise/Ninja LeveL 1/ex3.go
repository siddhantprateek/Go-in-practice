package main

import (
	"fmt"
)
var x int = 42
var y string = "James"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	g := fmt.Sprintln(x, y, z)
	//sprintf stands for "String print". Instead of printing on console, it store output on char buffer which are specified in sprintf.
	fmt.Println(s)
	fmt.Println(g)
}