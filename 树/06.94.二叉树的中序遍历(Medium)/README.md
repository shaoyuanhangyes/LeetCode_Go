# 6(94) 二叉树的中序遍历

## 描述

给定一个二叉树，返回它的中序遍历

## 示例

给定二叉树 [3,9,20,null,null,15,7]，

```bash
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
``` 

递归算法很简单，你可以通过迭代算法完成吗？

## Description

Given a binary tree, return the inorder traversal of its nodes' values.

## Example

Given binary tree [3,9,20,null,null,15,7]
```bash
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
```

Follow up: Recursive solution is trivial, could you do it iteratively?

`BitDance` `Amazon` `Microsoft` `Adobe` `Apple` `Google` `Tencent` `FaceBook` `Mi` `HuaWei` `Baidu` `DiDi` `Yahoo` `eBay` `Uber` `NetEase`

## 解题

```go
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
```
中序遍历基于栈的迭代算法的时间复杂度O(n) 空间复杂度O(n)