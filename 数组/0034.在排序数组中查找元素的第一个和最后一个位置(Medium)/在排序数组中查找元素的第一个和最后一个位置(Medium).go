package main

import "fmt"

func searchRange(nums []int, target int) []int {
	first := FirstPosition(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := LastPosition(nums, target)
	return []int{first, last}
}

func FirstPosition(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low < len(nums) && nums[low] == target {
		return low
	}
	return -1
}

func LastPosition(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	res := searchRange(nums, 8)
	fmt.Println(res)
}
