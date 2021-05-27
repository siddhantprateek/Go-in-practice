package main

import (
	"fmt"
)

func main(){

    var a int = 100
    var b int = 200

    var result int

    /* calling max function */
    result = max(a, b)

    fmt.Printf("Max: %d", result)
}

//  func <func-name>(arg) <return type>
func max(num1, num2 int) int {

    if (num1 > num2){
        return num1
    } else {
        return num2
    }

}
