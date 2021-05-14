# 8(144) 二叉树的前序遍历(Medium)

## 描述

给定一个二叉树，返回它的 前序 遍历

## 示例

```bash
输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [1,2,3]
``` 

进阶: 递归算法很简单，你可以通过迭代算法完成吗？

## Description

Given a binary tree, return the preorder traversal of its nodes' values.

## Example

```bash

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]

```
Follow up: Recursive solution is trivial, could you do it iteratively?

`BitDance` `Amazon` `Microsoft` `Adobe` `Apple` `Google` `Tencent` `Mi` `HuaWei` `VMware`

## 解题

### 迭代解法
```go
func preorder(root *TreeNode) []int {
	res := []int{}
	stack :=[]*TreeNode{}
	for len(stack)>0 || root != nil {
		for root !=nil {
			stack = append(stack,root) //进栈
			res = append(res,root.Val)
			root = root.Left
		}
		root = stack[len(stack)-1] //获取栈顶元素
		stack = stack[:len(stack)-1] //出栈
		root = root.Right
	}
	return res
}
```