package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//数组按序转化为二叉树
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

//前序遍历
func preorder(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

//显示一维数组内的所有元素 用于显示前序中序后序遍历的序列
func Showres(res []int) {
	fmt.Print("[")
	for _, i := range res {
		fmt.Print(i, " ")
	}
	fmt.Print("]")
}

func main() {
	nums := []int{7, 3, 5, 5, 6, 0, 8}
	root := createTree(nums, len(nums), 0)
	fmt.Print("前序遍历结果为")
	preorderlist := preorder(root)
	Showres(preorderlist)
}
