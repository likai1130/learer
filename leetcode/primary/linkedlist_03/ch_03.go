package main

import (
	"encoding/json"
	"fmt"
)

/**
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

*/
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead


}

func main() {
	lastNode2 := &ListNode{5, nil}
	lastNode := &ListNode{4, lastNode2}
	treeNode := &ListNode{3, lastNode}
	secondNode := &ListNode{2, treeNode}
	firstNode := &ListNode{1, secondNode}
	end := reverseList(firstNode)
	bytes, _ := json.Marshal(end)
	fmt.Println(string(bytes))

}