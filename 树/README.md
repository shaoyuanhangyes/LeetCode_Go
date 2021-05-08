## Golang模版


二叉树结构体
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

创建二叉树

```go
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
```


二叉树的层序遍历 输出到一个二维数组中

```go
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
```

打印二维数组

```go
func PrintMartrix(res [][]int) {
	for i := 0; i < len(res); i++ {
		fmt.Print("[")
		for j := 0; j < len(res[i]); j++ {
			fmt.Print(" ", res[i][j], " ")
		}
		fmt.Println("]")
	}
}
```