package main 
import "fmt"

func main(){
	var target int = 9
	nums := [8]int{7,5,2,8,9,4,1,6}
	var index int = -1
	for i := 0 ; i < 8 ; i ++ {
		if (nums[i] == target){
			index = i
			fmt.Println("Found at index", index)
		}
	}
	if (index == -1){
		fmt.Println("Not Found")
	}

}

