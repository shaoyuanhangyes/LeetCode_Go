package main

import (
	"fmt"
)

//链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

//使用数组创建对应的链表 后插法
func createList(nums []int) *ListNode {
	prev := new(ListNode)
	prev.Val = nums[0]
	prevHead := new(ListNode)
	prevHead = prev
	for i := 1; i < len(nums); i++ {
		temp := new(ListNode)
		temp.Val = nums[i]
		prev.Next = temp
		prev = temp
	}
	return prevHead
}

//显示链表中的所有元素
func ShowList(L *ListNode) {
	temp := L
	for temp != nil {
		fmt.Print(temp.Val, " ")
		temp = temp.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	l1 := createList(nums)
	l1 = reverseList(l1)
	ShowList(l1)
}
