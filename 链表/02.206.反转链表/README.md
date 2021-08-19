# 2(206) 反转链表
## 描述

反转一个单链表

## 示例

```bash
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
``` 

## Description

Reverse a singly linked list.

## Example

```bash
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
```

`BitDance` `Amazon` `Microsoft` `Alibaba` `FaceBook` `Google` `Tencent` `Adobe` `VIVO` `NVDIA` `Mi` `HuaWei` `Cisco` `Yahoo` `Uber` `NetEase` `Visa` `PayPal` `Lyft`

## 解题

```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
```