package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(nums []int, len int, index int) *TreeNode {
	root := new(TreeNode)
	if index < len && nums[index] != -1 {
		root.Val = nums[index]
		if 2*index+1 < len {
			root.Left = createTree(nums, len, 2*index+1)
		}
		if 2*index+2 < len {
			root.Right = createTree(nums, len, 2*index+2)
		}
	}
	return root
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

func Showres(res []int) {
	for _, i := range res {
		fmt.Print(i, " ")
	}
}

func main() {
	nums1 := []int{7, 3, 5, 5, 6, 0, 8}
	root := createTree(nums1, len(nums1), 0)
	res := inorderTraversal(root)
	Showres(res)
}
