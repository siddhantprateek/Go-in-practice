package main

import (
	"fmt"
)

func main() {
	n := "Q"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}

	fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}