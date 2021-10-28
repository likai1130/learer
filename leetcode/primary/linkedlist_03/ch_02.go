package main

import (
	"encoding/json"
	"fmt"
)

/**
	删除链表的倒数第N个节点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？
 */
type ListNode2 struct {
	Val int
	Next *ListNode2
}

func removeNthFromEnd(head *ListNode2, n int) *ListNode2 {
	//找到被删除节点的前继
	pre := head
	last := nodeSize(head) - n
	if (last == 0){
		return head.Next;
	}
	for i := 0; i < last-1; i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return head
}

func nodeSize(head *ListNode2) int {
	lenth := 0
	for  {
		if head != nil {
			lenth ++
			head = head.Next
		}else {
			return lenth
		}
	}
	return lenth
}

func main() {
	lastNode2 := &ListNode2{5, nil}
	lastNode := &ListNode2{4, lastNode2}
	treeNode := &ListNode2{3, lastNode}
	secondNode := &ListNode2{2, treeNode}
	firstNode := &ListNode2{1, secondNode}
	end := removeNthFromEnd(firstNode, 2)
	bytes, _ := json.Marshal(end)

	fmt.Println(string(bytes))

}
