package main

import (
	"encoding/json"
	"fmt"
)

/**
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。

现有一个链表 -- head = [4,5,1,9]，它可以表示为:

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnarn7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	     Val int
//	     Next *ListNode
//}

func deleteNode(node *ListNode,) {
	node.Val = node.Next.Val;
	node.Next = node.Next.Next
	bytes, _ := json.Marshal(node)
	fmt.Println(string(bytes))
}

func main() {
	lastNode := &ListNode{9, nil}
	treeNode := &ListNode{1, lastNode}
	secondNode := &ListNode{5, treeNode}
	firstNode := &ListNode{4, secondNode}
	bytes, _ := json.Marshal(firstNode)
	fmt.Println(string(bytes))
	deleteNode(firstNode)
}
