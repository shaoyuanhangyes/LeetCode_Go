package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return length
	}
	i := 0
	for j := 1; j < length; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}

func main() {
	nums := []int{1, 1, 2, 2, 2, 3}
	length := removeDuplicates(nums)
	fmt.Println(length)
	nums = nums[:length] 
	fmt.Println(nums)
}
