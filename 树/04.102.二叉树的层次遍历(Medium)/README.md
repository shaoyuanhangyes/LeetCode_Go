# 4(102) 二叉树的层次遍历

## 描述

给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）

## 示例

给定二叉树 [3,9,20,null,null,15,7]

```bash
    3
   / \
  9  20
    /  \
   15   7
``` 

返回其层次遍历结果：

```bash
[
  [3],
  [9,20],
  [15,7]
]
```

## Description

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

## Example

return its level order traversal as:

```bash
    3
   / \
  9  20
    /  \
   15   7
```

return its bottom-up level order traversal as:

```bash
[
  [3],
  [9,20],
  [15,7]
]
```

`BitDance` `Amazon` `Microsoft` `Adobe` `Apple` `Google` `Alibaba` `FaceBook` `Linkedln` `Tencent` `Baidu` `Cisco` `DiDi` `VMware` `Uber`

## 解题

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