package main

import "fmt"

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

func PrintMartrix(res [][]int) {
	for i := 0; i < len(res); i++ {
		fmt.Print("[")
		for j := 0; j < len(res[i]); j++ {
			fmt.Print(" ", res[i][j], " ")
		}
		fmt.Println("]")
	}
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		temp := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		queue = temp
	}
	return res
}

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := createTree(nums, len(nums), 0)
	res := levelOrder(root)
	PrintMartrix(res)
}
