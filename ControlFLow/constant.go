package main

import (
	"fmt"
)


func main() {

	const s = "Hello Siddhant"

	fmt.Println(s)
	fmt.Printf("%T\n", s)
	const b string = "siddhant"
	
	fmt.Printf("%T\n", b)
	bs := []byte(b)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	//  "%U" convert character into unicode or ASCII value
	for i := 0; i < len(s); i++{
		fmt.Printf("%#U ", s[i])

	}
}