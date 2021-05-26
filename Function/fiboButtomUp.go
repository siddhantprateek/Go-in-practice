package main
import "fmt"

func fiboBU(n int) (ret int) {
    // var f int
    f_minus_1, f_minus_2 := 1, 0
    for i := 2; i < n + 1; i++ {
        f := f_minus_1 + f_minus_2
        f_minus_2, f_minus_1 = f_minus_1, f
    }
    return f_minus_1

}


func main() {
    for i := 1; i < 20; i++{
        fmt.Printf("%d ",fiboBU(i))
    }
}
