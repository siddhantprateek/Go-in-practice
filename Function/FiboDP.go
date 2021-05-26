package main
import "fmt"


// func fiboDP(n int, dp []int)  int {
//
//     if n < 2{
//         return n
//     }
//     if dp[n]{
//         return dp[n]
//     }
//     dp[n] = fiboDP(n - 1, dp) + fiboDP(n - 2, dp)
//     return dp[n]
// }

func main () {

    // var n int
    // fmt.Printf("Enter value to be printed in FiboSeries:")
    // fmt.Scan(&n)

    /*
    I assume go does not enjoy users instantiating an Array's length with a
    calculated value. It only accepts constants. Should I just give up and use a slice instead?
    I expect a slice is a dynamic array meaning it's either a linked list or
    copies into a larger array when it gets full.
    */

    var n int = 5
    var dp [10]int

    for i := 0;  i < 10 ; i++{
        dp[i] = 0
    }
    for i := 0;  i < n; i++{
        fmt.Printf("%d ", dp[i])
    }

    // for i := 0; i < n; i++ {
    //     fmt.Println(fiboDP(i, dp))
    // }
}
