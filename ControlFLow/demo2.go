package main

import "fmt"

func main() {
	fmt.Println("Ka boom")

	foo()

	fmt.Println("Iteration")

	for i := 0; i < 100; i++ {
		if i % 2 == 0{
			fmt.Printf("%d ",i)
		}
	}
}

func foo() {
	fmt.Printf("I'm foo")
}