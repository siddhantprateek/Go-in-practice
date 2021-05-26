package main

import "fmt"

func main() {

	nums := [8]int{17,5,12,8,9,4,11,6}
	for i := 0; i < len(nums); i++{
		fmt.Printf("%d ",nums[i])
	}
	fmt.Printf("\n")

	for i := 0; i < len(nums); i++{
		for j := 0; j < len(nums) - i - 1; j++{
			if nums[j] > nums[j + 1]{
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}
	fmt.Printf("Sorted:")
	for i := 0; i < len(nums); i++{
		fmt.Printf("%d ",nums[i])
	}

}