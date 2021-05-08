# 2(21) 合并两个有序链表

## 描述

将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

## 示例

```bash
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
``` 

## Description

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

## Example

```bash
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

```
`BitDance` `Amazon` `Microsoft` `Alibaba` `FaceBook` `Google` `Tencent` `Apple` `HuaWei` `Baidu` `Adobe` `Intel` `VMware` `Uber` `Yahoo`


## 解题

Golang版本

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := new(ListNode)
	prev.Val = -1
	prevHead := prev
	for l1!=nil&&l2!=nil {
		if l1.Val<l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else{
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1!=nil {
		prev.Next = l1
	}else{
		prev.Next = l2
	}
	return prevHead.Next
}
```