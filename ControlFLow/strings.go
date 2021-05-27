package main

import (
	"fmt"
)


func main() {

	s := "Hello Siddhant"

	fmt.Println(s)
	s = "AHello Bro"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	//  "%U" convert character into unicode or ASCII value
	for i := 0; i < len(s); i++{
		fmt.Println("%#U", s[i])

	}

	fmt.Println("")
	// %x : Hexadecimal
	// "%#x" gives the hex value of the no. or character 

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
	fmt.Println("--------------------")
	g := "Hello, playground"
	fmt.Println(g)
	fmt.Printf("%s\n", g)
	fmt.Printf("%q\n", g)
	fmt.Printf("-> %x\n", g)

	for i := 0; i < len(g); i++ {
		fmt.Printf(" %x ", g[i])
	}
	fmt.Println("")
	for i := 0; i < len(g); i++ {
		fmt.Printf("%d ", g[i])
	}
	fmt.Println("")

	// %#U gives unicode with the character itself
	for i, v := range g {
		fmt.Printf("%#U -> %d \n", v, i)
	}
}